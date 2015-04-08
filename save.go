package main

import (
	`encoding/json`
	`flag`
	`log`
	`os`
	`time`
)

var (
	saveFile string
)

func init() {
	flag.StringVar(&saveFile, `backup`, `backup.json`, `json file for backup`)
}

func saveChallenge() {
	for {
		file, err := os.Create(saveFile)
		if err != nil {
			log.Fatalf(`unable to open file '%s' for backup`, saveFile)
		}

		jsonChallenge.RLock()
		stats := jsonChallenge
		jsonChallenge.RUnlock()

		data, _ := json.Marshal(stats)
		file.Write(data)

		file.Close()
		time.Sleep(30 * time.Second)
	}
}
