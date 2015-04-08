package main

import ()

func NewJSONStages() map[int]stage {
	data := make(map[int]stage, 3)

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
