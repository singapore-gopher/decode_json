package main

import (
//`encoding/json`
//`io/ioutil`
//`log`
//`os`
)

var (
	jsonChallenge challenge
)

func init() {
	jsonChallenge = challenge{
		Teams: make(map[string]map[string]stageStats),
	}
}

func setupChallenge() {
	var stages map[string]stage

	// file, err := os.Open(saveFile)
	// if os.IsNotExist(err) { // no file there
	// 	stages = NewJSONStages()
	// } else {
	// 	data, err := ioutil.ReadAll(file)
	// 	if err != nil {
	// 		log.Fatalf(`could not read existing save; err=%v`, err)
	// 	}
	// 	if err := json.Unmarshal(data, &stages); err != nil {
	// 		log.Fatalf(`could not decode save; err=%v`, err)
	// 	}
	// 	file.Close()
	// }

	stages = NewJSONStages()
	jsonChallenge.Lock()
	jsonChallenge.Stages = stages
	jsonChallenge.Unlock()
}

func NewJSONStages() map[string]stage {
	data := make(map[string]stage, 3)

	data[Stage1] = stage{
		Tests: []testCase{
			{`{"first":5,"second":10}`, `{"sum":15}`},
			{`{"first":7,"second":234}`, `{"sum":241}`},
			{`{"first":9,"second":8}`, `{"sum":17}`},
			{`{"first":14,"second":84}`, `{"sum":98}`},
		},
	}

	return data
}
