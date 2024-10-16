package lib

import (
	"errors"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

var Version = "0.0.1"

func IsWaylandDisplayManager() bool {
	envVars := []string{"XDG_SESSION_TYPE", "XDG_BACKEND", "QT_QPA_PLATFORM", "WAYLAND_DISPLAY"}

	for _, env := range envVars {
		val := os.Getenv(env)
		if strings.Contains(val, "wayland") {
			return true
		}
	}

	return false
}

// Identify relevant display managers (Linux).
func HasWaylandDisplayManager() bool {
	if runtime.GOOS != "linux" {
		return false
	}

	return IsWaylandDisplayManager()
}

/*
Check if wlroots compositor is running, as grim only supports wlroots.
Certainly not wlroots based are: KDE, GNOME and Unity.
Others are likely wlroots based.
*/
func HasWLRootsCompositor() bool {
	if runtime.GOOS != "linux" {
		return false
	}

	kde_full_session := strings.ToLower(os.Getenv("KDE_FULL_SESSION"))
	xdg_current_desktop := strings.ToLower(os.Getenv("XDG_CURRENT_DESKTOP"))
	desktop_session := strings.ToLower(os.Getenv("DESKTOP_SESSION"))
	gnome_desktop_session_id := strings.ToLower(os.Getenv("GNOME_DESKTOP_SESSION_ID"))

	if gnome_desktop_session_id == "this-is-deprecated" {
		gnome_desktop_session_id = ""
	}

	if gnome_desktop_session_id != "" || strings.Contains(xdg_current_desktop, "gnome") || kde_full_session != "" || strings.Contains(desktop_session, "kde-plasma") || strings.Contains(xdg_current_desktop, "unity") {
		return false
	}

	return true
}

/*
Detect Gnome version of current session.
Returns: Version string or empty string if not detected.
*/
func GetGnomeVersion() (string, error) {
	if runtime.GOOS != "linux" {
		err := errors.New("Unsupported platform. Gnome shell not available")
		return "", err
	}

	if os.Getenv("GNOME_DESKTOP_SESSION_ID") == "" && !strings.Contains(strings.ToLower(os.Getenv("XDG_CURRENT_DESKTOP")), "gnome") {
		err := errors.New("GNOME_DESKTOP_SESSION_ID/XDG_CURRENT_DESKTOP env not found")
		return "", err
	}

	which_cmd := exec.Command("which", "gnome-shell")
	_, err := which_cmd.Output()
	if err != nil {
		LogError.Println("Failed to run cmd 'which'")
		return "", err
	}

	gnome_version_cmd := exec.Command("gnome-shell", "--version")
	gnome_version_output, err := gnome_version_cmd.Output()
	if err != nil {
		LogWarning.Fatalf("Error when trying to get gnome version: %v", err)
		return "", err
	}

	re, err := regexp.Compile(`\s+([\d\.]+)`)
	if err != nil {
		LogError.Println("Error while compiling regex")
		return "", err
	}
	result := strings.Trim(re.FindString(strings.Trim(string(gnome_version_output), " ")), " ")
	return result, nil
}

func IsFlatpakPackage() bool {
	return os.Getenv("FLATPAK_ID") != ""
}

func IsAppImagePackage() bool {
	return os.Getenv("APPIMAGE") != ""
}
