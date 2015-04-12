package main

import (
	`encoding/json`
	`errors`
	`fmt`
	`io/ioutil`
	`net/http`
)

type registerInput struct {
	Name string `json:"name"`
}

var (
	ErrNeedTeamName  = errors.New(`need to have a name`)
	ErrDuplicateName = errors.New(`name already registered`)
)

func register(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		NewError(resp, http.StatusBadRequest, err)
		return
	}

	var input registerInput
	if err := json.Unmarshal(body, &input); err != nil {
		NewError(resp, http.StatusBadRequest, err)
		return
	}

	if len(input.Name) < 1 {
		NewError(resp, http.StatusBadRequest, ErrNeedTeamName)
		return
	}

	jsonChallenge.Lock()
	if _, exist := jsonChallenge.Teams[input.Name]; exist {
		jsonChallenge.Unlock()
		NewError(resp, http.StatusBadRequest, ErrDuplicateName)
		return
	}

	jsonChallenge.Teams[input.Name] = make(map[string]stageStats)
	jsonChallenge.Unlock()

	resp.WriteHeader(http.StatusCreated)
	resp.Write([]byte(fmt.Sprintf(`{"ok":true,"name":"%s"}`, input.Name)))
	resp.Write([]byte{'\n'})
}
