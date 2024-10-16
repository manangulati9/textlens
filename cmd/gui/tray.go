package gui

import (
	"textlens/internal/lib"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

const (
	_EXIT_DELAY            = 5 // In seconds
	_UPDATE_CHECK_INTERVAL = 7 // In days
)

var (
	_NORMAL_ICON *gui.QIcon
	_DONE_ICON   *gui.QIcon
)

func NewSystemTray(app *widgets.QApplication, args *lib.Args) *widgets.QSystemTrayIcon {
	_NORMAL_ICON = gui.NewQIcon5("resources/icons/tray.svg")
	_DONE_ICON = gui.NewQIcon5("resources/icons/tray_done.svg")
	tray := widgets.NewQSystemTrayIcon2(_NORMAL_ICON, app)
	return tray
}
