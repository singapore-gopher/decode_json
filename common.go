package main

import (
	`errors`
	`fmt`
	`net/http`
	`strings`
	`sync`
	`time`
)

// testCase stores the info of one test
type testCase struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

// stage stores the info of a challege
type stage struct {
	Tests []testCase `json:"tests"`
}

// challenge stores the whole 3-stage challenge
type challenge struct {
	sync.RWMutex
	Stages map[string]stage                 `json:"stages"`
	Teams  map[string]map[string]stageStats `json:"teams"`
}

type stageStats struct {
	Attempts      int       `json:"attempts"`
	Passed        int       `json:"passed"`
	FirstAttempt  time.Time `json:"first_try"`
	LatestAttempt time.Time `json:"latest_try"`
}

const (
	Stage1 = `stage_1` //make it easier to encode leaderboard
	Stage2 = `stage_2`
	Stage3 = `stage_3`
)

// common error across stages
var (
	ErrNotEnoughData = errors.New(`number of solutions != number of test cases`)
	ErrNeedTeamInfo  = errors.New(`no team info included`)
	ErrUnknownTeam   = errors.New(`unknown team`)
)

func NewError(resp http.ResponseWriter, status int, err error) {
	resp.WriteHeader(status)
	resp.Write([]byte(`{"error":"` + err.Error() + `"}`))
	resp.Write([]byte{'\n'})
}

func updateTeam(name, stage string, passed int) {
	jsonChallenge.Lock()
	now := time.Now()
	if _, retry := jsonChallenge.Teams[name][stage]; !retry {
		jsonChallenge.Teams[name][stage] = stageStats{
			Attempts:      0,
			Passed:        0,
			FirstAttempt:  now,
			LatestAttempt: now,
		}
	}

	stats := jsonChallenge.Teams[name][stage]
	stats.Attempts++
	if passed > stats.Passed {
		stats.Passed = passed
	}
	stats.LatestAttempt = now
	jsonChallenge.Teams[name][stage] = stats
	jsonChallenge.Unlock()
}

func dataHandler(stage string) func(resp http.ResponseWriter, req *http.Request) {
	jsonChallenge.RLock()
	stages := jsonChallenge.Stages
	jsonChallenge.RUnlock()

	var data []string
	for _, test := range stages[stage].Tests {
		data = append(data, test.Input)
	}
	results := fmt.Sprintf(`{"inputs":[%s]}`, strings.Join(data, `,`))

	return func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte(results))
	}
}
