package handlers

import "github.com/therecipe/qt/gui"

type Handler interface {
	IsCompatible() bool
	IsInstalled() bool
	Capture() ([]*gui.QImage, error)
}
