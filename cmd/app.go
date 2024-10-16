package cmd

import (
	"os"
	"textlens/cmd/gui"
	"textlens/internal/lib"

	"github.com/therecipe/qt/widgets"
)

/*
Parse command line arguments.
Exit if Textlens was started with --version flag.
Auto-enable tray for "background mode", which starts Textlens in tray without
immediately opening the select-region window.
*/
func GetArgs() *lib.Args {
	args := lib.ParseArgs()
	args.ValidateArgs()
	return args
}

func Prepare() (*widgets.QApplication, *widgets.QSystemTrayIcon) {
	args := GetArgs()

	PrepareLogging(args.Verbosity)
	PrepareEnvs()

	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetQuitOnLastWindowClosed(false)

	tray := gui.NewSystemTray(app, args)
	return app, tray
}

/*
Prepare environment variables depending on setup and system.
*/
func PrepareEnvs() {
	switch {
	case lib.IsWaylandDisplayManager():
		lib.SetWaylandEnvs()
	case lib.IsFlatpakPackage():
		lib.SetFlatpakEnvs()
	case lib.IsAppImagePackage():
		lib.SetAppImageEnvs()
	}
}

func PrepareLogging(log_level string) {
	lib.NewLogger(log_level)
	lib.LogInfo.Printf("Start Textlens v%s\n", lib.Version)
	lib.LogInfo.Println("Preparing app...")
}
