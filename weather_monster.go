package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sharangj/weather_monster/config"
	"github.com/sharangj/weather_monster/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: weather_monster -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	r := server.Init()

	r.Run() // listen and serve on 0.0.0.0:8080
}
