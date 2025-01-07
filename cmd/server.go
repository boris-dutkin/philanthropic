package main

import (
	"github.com/boris-dutkin/philanthropic/config"
	"github.com/boris-dutkin/philanthropic/database"
)

func main() {
	config := config.New()
	database.Connect(config)
}
