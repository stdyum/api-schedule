package main

import (
	"log"

	"github.com/stdyum/api-schedule/internal"
)

func main() {
	log.Fatalf("error launching web server %s", internal.App().Run())
}
