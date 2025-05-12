package main

import (
	"github.com/joho/godotenv"
	"github.com/zuhrulumam/doit-test/cmd"
)

func main() {
	_ = godotenv.Load()
	cmd.Execute()
}
