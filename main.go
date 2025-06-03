package main

import (
	"github.com/joho/godotenv"
	"github.com/zuhrulumam/go-parking-lot/cmd"
)

func main() {
	_ = godotenv.Load()
	cmd.Execute()
}
