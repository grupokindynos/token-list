package main

import (
	"encoding/json"
	"github.com/grupokindynos/token-list/cmd"
	"io/ioutil"
)

func main() {
	raw, err := json.Marshal(cmd.ListData)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./kindynos.tokens.json", raw, 0700)
	if err != nil {
		panic(err)
	}
}
