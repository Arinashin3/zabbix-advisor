package main

import (
	"gioui.org/app"
	"gioui.org/unit"
	"github.com/arinashin3/zabbix-advisor/src/draws"
	"log"
	"os"
)

func main() {
	go func() {
		// create new window
		w := app.NewWindow(
			app.Title("Zabbix-Advisor"),
			app.Size(unit.Dp(1000), unit.Dp(600)),
		)
		if err := draws.Draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
