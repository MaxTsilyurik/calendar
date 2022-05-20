package main

import "calendar/internal/runner"

const configDirectory = "configs"

func main() {
	runner.Start(configDirectory)
}
