package main
import (
	app "Week04/app"
	server "Week04/server"
	"os"
)

func main() {
	myApp := app.New()
	myServer, err := server.New("127.0.0.1:9000")
	if err != nil{
		os.Exit(1)
	}
	newHook := app.Hook{
		OnStart:myServer.OnStart,
		OnStop:myServer.OnStop,
	}
	myApp.AppendHook(newHook)
	myApp.Run()
}