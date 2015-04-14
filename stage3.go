package main

import (
	"encoding/json"
	"fmt"
	"github.com/singapore-gophers/decode_json/stream"
	"net/http"
)

type stage3request struct {
	TeamName string   `json:"team"`
	Values   []string `json:"values"`
}

type stage3response struct {
	Correct int    `json:"correct"`
	Result  string `json:"result"`
}

func stage3SubmitHandler(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	solution := &stage3request{}

	err := dec.Decode(solution)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	goodValues := verifyStage3Solution(solution)

	if goodValues == 3 {
		updateTeam(solution.TeamName, "stage3", 1)
	}

	w.Header().Set("Content-Type", "application/json")
	response := &stage3response{
		Correct: goodValues,
		Result: fmt.Sprintf("%v/3 %s", goodValues, map[int]string{
			0: "Don't give up!",
			1: "At least you got that one right, keep up!",
			2: "Almost there!",
			3: "You made it! Awesome!",
		}[goodValues]),
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// 0 - No correct value
// 1 - 3 Correct values
func verifyStage3Solution(s *stage3request) int {
	correct := 0
	for _, val := range s.Values {
		if stream.IsMagicValue(val) {
			if correct < 3 {
				correct++
			}
		}
	}

	return correct
}
