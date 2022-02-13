package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	InitialBalance float64 `json:"initialBalanceAmount"`
	MinumumBalance float64 `json:"minimumBalanceAmount"`
}

var c = &Config{}

func init() {
	/*
		file, err := os.Open(".config/local.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		read, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(read, c)
		if err != nil {
			panic(err)
		}
	*/
}

func Get() *Config {
	// read as file
	file, err := os.Open(".config/local.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// convert to byte[]
	read, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// convert byte[] to struct
	err = json.Unmarshal(read, c)
	if err != nil {
		panic(err)
	}

	return c
}
