package cmd

import (
	"os"
	"textlens/cmd/screengrab/handlers"
	"textlens/internal/lib"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func Prepare() (app *widgets.QApplication, tray *widgets.QSystemTrayIcon) {
	lib.NewLogger()
	lib.Logger.Info.Println("Logging initialized")
	lib.Logger.Info.Println("Preparing app...")

	PrepareEnvs()

	app = widgets.NewQApplication(len(os.Args), os.Args)
	app.SetQuitOnLastWindowClosed(false)

	tray = widgets.NewQSystemTrayIcon2(gui.NewQIcon5("resources/icons/tray.svg"), app)
	_, err := handlers.NewGrimHandler().Capture()
	if err != nil {
		lib.Logger.Error.Fatalln("Grim capture errored")
	}
	return app, tray
}

/*
Prepare environment variables depending on setup and system.
Enable exiting via CTRL+C in Terminal.
*/
func PrepareEnvs() {
	// Allow closing QT app with CTRL+C in terminal
	// signal.signal(signal.SIGINT, signal.SIG_DFL)

	if lib.IsWaylandDisplayManager() {
		lib.SetWaylandEnvs()
	}
	// if system_info.is_flatpak_package():
	//     utils.set_environ_for_flatpak()
	// if system_info.is_appimage_package():
	//     utils.set_environ_for_appimage()
}
