package main

import (
	`encoding/json`
	`io/ioutil`
	`math/rand`
	`net/http`
	`strings`
	`time`
)

type stage2Input struct {
	TeamName string `json:"team"`
	Faulty   []int  `json:"faulty"`
}

const (
	faultyPerTurn = 2
	goodPerTurn   = 3
)

var (
	stage2Faulty []string = []string{
		`{"id":1,     "temperature":25, "humidity": 40}`,
		`{"id":163,   "have_smoke":false, "humidity": 40}`,
		`{"id":569,   "have_smoke":false, "humidity": 40}`,
		`{"id":1033,  "have_smoke":false, "temperature":25}`,
		`{"id":73,    "temperature":25,   "humidity": 40}`,
		`{"id":2141,  "temperature":25,   "humidity": 40}`,
		`{"id":19,    "have_smoke":false, "temperature":25}`,
		`{"id":419,   "temperature":25,   "humidity": 40}`,
		`{"id":1549,  "have_smoke":false, "humidity": 40}`,
		`{"id":1621,  "have_smoke":false, "temperature":25}`,
	}
	stage2Good []string = []string{
		`{"id":2,  "have_smoke":false, "temperature":25, "humidity": 40}`,
		`{"id":4,  "have_smoke":false, "temperature":25, "humidity": 40}`,
		`{"id":6,  "have_smoke":false, "temperature":25, "humidity": 40}`,
		`{"id":8,  "have_smoke":false, "temperature":25, "humidity": 40}`,
		`{"id":10, "have_smoke":false, "temperature":25, "humidity": 40}`,
		`{"id":12, "have_smoke":false, "temperature":25, "humidity": 40}`,
		`{"id":14, "have_smoke":false, "temperature":25, "humidity": 40}`,
		`{"id":16, "have_smoke":false, "temperature":25, "humidity": 40}`,
		`{"id":18, "have_smoke":false, "temperature":25, "humidity": 40}`,
		`{"id":20, "have_smoke":false, "temperature":25, "humidity": 40}`,
	}
)

func init() {
	rand.Seed(time.Now().Unix())
}

func stage2DataHandler(resp http.ResponseWriter, req *http.Request) {
	faulty, good := make([]string, 0, 2), make([]string, 0, 3)
	usedFaulty, usedGood := make(map[int]bool), make(map[int]bool)

	for len(faulty) < faultyPerTurn {
		tmp := rand.Intn(10)
		if _, exist := usedFaulty[tmp]; exist {
			continue
		}
		faulty = append(faulty, stage2Faulty[tmp])
		usedFaulty[tmp] = true
	}
	for len(good) < goodPerTurn {
		tmp := rand.Intn(10)
		if _, exist := usedGood[tmp]; exist {
			continue
		}
		good = append(good, stage2Good[tmp])
		usedGood[tmp] = true
	}
	result := `{"sensors":[` + strings.Join(faulty, `,`) + `,` + strings.Join(good, `,`) + `]}`

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(result))
}

func stage2SubmitHandler(resp http.ResponseWriter, req *http.Request) {
	// parse input
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		NewError(resp, http.StatusBadRequest, err)
		return
	}

	var input stage2Input
	if err := json.Unmarshal(body, &input); err != nil {
		NewError(resp, http.StatusBadRequest, err)
		return
	}

	if len(input.TeamName) < 1 {
		NewError(resp, http.StatusBadRequest, ErrNeedTeamInfo)
		return
	}

	// check tests
	jsonChallenge.RLock()
	_, teamExists := jsonChallenge.Teams[input.TeamName]
	jsonChallenge.RUnlock()

	if !teamExists {
		NewError(resp, http.StatusBadRequest, ErrUnknownTeam)
		return
	}

	var correct int = 0
	for _, id := range input.Faulty {
		if id == 1 || id == 163 || id == 569 || id == 1033 || id == 73 || id == 2141 || id == 19 || id == 419 || id == 1549 || id == 1621 {
			correct++
		}
	}

	resp.WriteHeader(http.StatusOK)
	if correct == faultyPerTurn {
		updateTeam(input.TeamName, Stage2, 1)
		resp.Write([]byte(`{"ok":true}`))
		return
	}

	updateTeam(input.TeamName, Stage2, 0)
	resp.Write([]byte(`{"error":"try again"`))
}
