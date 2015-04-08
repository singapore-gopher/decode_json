package main

import (
	`encoding/json`
	`fmt`
	`io/ioutil`
	`net/http`
)

type stage1Out struct {
	Sum int `json:"sum"`
}

type stage1Input struct {
	TeamName  string      `json:"team"`
	Solutions []stage1Out `json:"solutions"`
}

func stage1Handler(resp http.ResponseWriter, req *http.Request) {
	// parse input
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		NewError(resp, http.StatusBadRequest, err)
		return
	}

	var input stage1Input
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
	stage1Count := len(jsonChallenge.Stages[Stage1].Tests)
	tests := jsonChallenge.Stages[Stage1].Tests
	_, teamExists := jsonChallenge.Teams[input.TeamName]
	jsonChallenge.RUnlock()

	if !teamExists {
		NewError(resp, http.StatusBadRequest, ErrUnknownTeam)
		return
	}

	if len(input.Solutions) != stage1Count {
		NewError(resp, http.StatusBadRequest, ErrNotEnoughData)
		return
	}

	var correct int = 0
	for i, test := range tests {
		guess := input.Solutions[i]
		if stage1Check(test, guess) {
			correct++
		}
	}

	// write out
	updateTeam(input.TeamName, Stage1, correct)
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(fmt.Sprintf(`{"passed":%d}`, correct)))
}

func stage1Check(test testCase, guess stage1Out) bool {
	var out stage1Out
	json.Unmarshal([]byte(test.Output), &out) // should be correct

	return guess.Sum == out.Sum
}
