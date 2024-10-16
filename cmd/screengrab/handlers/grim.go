package handlers

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"textlens/cmd/screengrab"
	"textlens/internal/lib"

	"github.com/therecipe/qt/gui"
)

type GrimHandler struct {
	InstallInstructions string
}

func NewGrimHandler() *GrimHandler {
	g := GrimHandler{
		"Install the package `grim` using your system's package manager.",
	}
	return &g
}

func (g *GrimHandler) IsCompatible() bool {
	return lib.HasWaylandDisplayManager() && lib.HasWLRootsCompositor()
}

func (g *GrimHandler) IsInstalled() bool {
	cmd := exec.Command("which", "grim")
	res, err := cmd.Output()
	if err != nil {
		lib.LogError.Fatalln("Failed to run command 'which'")
		return false
	}
	return string(res) != ""
}

/*
Capture screenshot with the grim CLI tool for wayland.
Is supported by some wayland compositors, e.g. hyprland. Won't work on e.g. standard Gnome.
*/
func (g *GrimHandler) Capture() ([]*gui.QImage, error) {
	const FILE_NAME = "textlens_grim_screenshot.png"
	path := os.TempDir() + `/` + FILE_NAME

	if strings.Contains(runtime.GOOS, "windows") {
		path = os.TempDir() + `\` + FILE_NAME
	}

	defer os.Remove(path)

	cmd := exec.Command("grim", path)
	_, err := cmd.Output()
	if err != nil {
		lib.LogError.Fatalln("Failed to catpure screenshot using 'grim'")
		return nil, err
	}

	full_image := gui.NewQImage9(path, "PNG")
	images, err := screengrab.SplitFullDesktopToScreens(full_image)
	if err != nil {
		return nil, err
	}

	return images, nil
}
