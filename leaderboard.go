package main

import (
	`sync`
)

// testCase stores the info of one test
type testCase struct {
	Input  string
	Output string
}

// stage stores the info of a challege
type stage struct {
	Tests []testCase
}

// challenge stores the whole 3-stage challenge
type challenge struct {
	sync.RWMutex
	Stages map[int]stage
	Teams  map[string]map[int]stageStats
}

type stageStats struct {
	Attempts int
	Passed   int
}

const (
	Stage1 = iota
	Stage2
	Stage3
)

var (
	jsonChallenge challenge
)

func init() {
	jsonChallenge = challenge{
		Teams: make(map[string]map[int]stageStats),
	}
}

func setupChallenge() {
	jsonChallenge.Lock()
	jsonChallenge.Stages = NewJSONStages()
	jsonChallenge.Unlock()
}
