package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	InitialBalance int     `json:"initialBalanceAmount"`
	MinumumBalance float64 `json:"minimumBalanceAmount"`
}

var C = &Config{}

func init() {
	file, err := os.Open(".config/local.json")
	handleErr(err)
	defer file.Close()

	read, err := io.ReadAll(file)
	handleErr(err)

	err = json.Unmarshal(read, C)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
