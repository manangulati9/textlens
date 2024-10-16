package gui

import "github.com/therecipe/qt/core"

type Communicate struct {
	core.QObject

	ExitApplication     func(float32)
	OnCopiedToClipboard func()
	OnImageCropped      func()
	OnRegionSelected    func(Rect)
	OnLanguagesChanged  func([]string)
}
