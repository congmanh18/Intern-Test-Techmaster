package main

import (
	"03/server"
	"flag"
	"log"
)

// @title CMN Express API Documentation
// @host 203.145.47.225
// @BasePath /
func main() {
	// Dùng cách cờ để đọc file config
	// go run ./cmd/main.go -config="./conf/pgsql.env"
	configPathFlag := flag.String("config", "", "Path to configuration file")
	flag.Parse()
	// export CONFIG_PATH="./conf/pgsql.env"
	// go run main.go
	configPath := *configPathFlag

	log.Printf("Starting server with config: %s", configPath)

	server.Run(configPath)
}
