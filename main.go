package main

import (
	"os"
	"textlens/cmd"
)

func main() {
	app, tray := cmd.Prepare()
	tray.Show()
	os.Exit(app.Exec())
}
