package handlers

import (
	"runtime"
	"textlens/internal/lib"

	"github.com/therecipe/qt/gui"
)

type QtHandler struct {
	InstallInstructions string
}

func NewQtHandler() *QtHandler {
	qt := QtHandler{}
	return &qt
}

func (qt *QtHandler) IsInstalled() bool {
	return true
}

func (qt *QtHandler) IsCompatible() bool {
	return runtime.GOOS == "linux" && !lib.HasWaylandDisplayManager()
}

func (qt *QtHandler) Capture() ([]*gui.QImage, error) {
	/* Capture screenshot with QT method and Screen object.
	   Works well on X11, fails on multi monitor macOS.
	*/
	var images []*gui.QImage
	for _, screen := range gui.QGuiApplication_Screens() {
		screenshot := screen.GrabWindow(0, 0, 0, -1, -1)
		image := screenshot.ToImage()
		images = append(images, image)
	}
	return images, nil
}
