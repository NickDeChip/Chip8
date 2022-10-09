package main

import (
	"encoding/hex"
	"os"
)

func main() {
	filePathArg := os.Args[1]
	file, err := os.ReadFile(filePathArg)
	if err != nil {
		panic(err)
	}
	prog, err := hex.DecodeString(string(file))
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("a", prog, 0644)
	if err != nil {
		panic(err)
	}

}
