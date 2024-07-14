package main

import (
	"vblog/cmd"

	_ "vblog/app"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
