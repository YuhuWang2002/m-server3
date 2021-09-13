package main

import (
	"github.com/mix-go/dotenv"
	"m-server3/commands"
	
	
	_ "m-server3/di"
	"github.com/mix-go/xcli"
)

func main() {
	xcli.SetName("app").
		SetVersion("0.0.0-alpha").
		SetDebug(dotenv.Getenv("APP_DEBUG").Bool(false))
	xcli.AddCommand(commands.Commands...).Run()
}
