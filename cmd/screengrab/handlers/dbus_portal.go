package handlers

import (
	"runtime"
	"strings"
	"textlens/internal/lib"

	"github.com/therecipe/qt/gui"
)

type DBusPortalHandler struct {
	InstallInstructions string
}

func NewDBusPortalHandler() *DBusPortalHandler {
	bus := DBusPortalHandler{}
	return &bus
}

func (dbp *DBusPortalHandler) IsCompatible() bool {
	// TODO: Update and check whether wayland allows for native screenshot capture using dbus_portal
	return runtime.GOOS == "linux" && !lib.HasWaylandDisplayManager()
}

func (dbp *DBusPortalHandler) IsInstalled() bool {
	version, err := lib.GetGnomeVersion()
	if err != nil {
		return true
	}
	parts := strings.Split(version, "")
	return parts[0] == "4" && parts[1] > "1"
}

func (dbp *DBusPortalHandler) Capture([]*gui.QImage, error) {
	/* Capture screenshots for all screens using org.freedesktop.portal.Desktop.
	   This methods works gnome-shell >=v41 and wayland.

	   In newer xdg-portal implementations, the first request has to be done in
	   "interactive" mode, before the application is allowed to query screenshots without
	   the dialog window in between.

	   As there is no way to query for that permission, we try both:
	   1. Try none-interactive mode
	   2. If timeout triggers, retry in interactive mode with a helper window */
}
