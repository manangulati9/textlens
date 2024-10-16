package handlers

import (
	"errors"
	"os"
	"runtime"
	"strings"
	"textlens/cmd/screengrab"
	"textlens/internal/lib"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/dbus"
	"github.com/therecipe/qt/gui"
)

type DBusShellHandler struct {
	InstallInstructions string
}

func NewDbusShellHandler() *DBusShellHandler {
	sh := DBusShellHandler{}
	return &sh
}

func (sh *DBusShellHandler) IsInstalled() bool {
	return true
}

func (sh *DBusShellHandler) IsCompatible() bool {
	return true
}

func (sh *DBusShellHandler) Capture() ([]*gui.QImage, error) {
	/* Capture screenshots for all screens using org.gnome.Shell.Screenshot.
	This methods works gnome-shell < v41 and wayland. */
	const ITEM = "org.gnome.Shell.Screenshot"
	const INTERFACE = "org.gnome.Shell.Screenshot"
	const PATH = "/org/gnome/Shell/Screenshot"

	bus := dbus.QDBusConnection_SessionBus().SessionBus()
	if !bus.IsConnected() {
		lib.LogError.Fatalln("Not connected to dbus!")
	}

	// Creating a new screenshot interface
	screenshot_interface := dbus.NewQDBusInterface2(ITEM, PATH, INTERFACE, nil, nil)
	if !screenshot_interface.IsValid() {
		msg := "Invalid dbus_interface"
		lib.LogError.Println(msg)
		return nil, errors.New(msg)
	}

	// Creating a temp file to
	const FILE_NAME = "textlens_grim_screenshot.png"
	path := os.TempDir() + `/` + FILE_NAME

	if strings.Contains(runtime.GOOS, "windows") {
		path = os.TempDir() + `\` + FILE_NAME
	}

	defer os.Remove(path)
	result := screenshot_interface.Call("Screenshot", core.NewQVariant9(true), core.NewQVariant9(false), core.NewQVariant12(path), nil, nil, nil, nil, nil)

	if result.ErrorName() != "" {
		lib.LogError.Println("Failed to capture screenshot using dbus_shell")
		return nil, errors.New(result.ErrorMessage())
	}

	image := gui.NewQImage9(path, "PNG")
	return screengrab.SplitFullDesktopToScreens(image), nil
}
