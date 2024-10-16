package screengrab

import (
	"errors"
	"textlens/internal/lib"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

/*
Split full desktop image into list of images per screen.
Also resizes screens according to image:virtual-geometry ratio.
*/
func SplitFullDesktopToScreens(full_image *gui.QImage) ([]*gui.QImage, error) {
	if full_image == nil {
		lib.LogError.Fatalln("Invalid image ptr provided. Exiting...")
		return nil, errors.New("Nil image ptr provided")
	}

	virtual_geometry := gui.QGuiApplication_PrimaryScreen().VirtualGeometry()
	ratio := full_image.Rect().Width() / virtual_geometry.Width()

	lib.LogInfo.Printf("Virtual geometry width: %d", virtual_geometry.Width())
	lib.LogInfo.Printf("Image width: %d", full_image.Rect().Width())
	lib.LogInfo.Printf("Resize ratio: %d", ratio)

	var images []*gui.QImage
	screen_cnt := widgets.QApplication_Desktop().NumScreens()
	if screen_cnt <= 1 {
		images = append(images, full_image)
		return images, nil
	}

	// QGuiApplication_Screens() errors out on my single monitor screen. Need to test for a multi-monitor setup
	screens, err := lib.RunUnsafeFunc(gui.QGuiApplication_Screens)
	if err != nil {
		return nil, err
	}

	for _, screen := range screens {
		geo := screen.Geometry()
		region := core.NewQRect()
		region.SetLeft(geo.X() * ratio)
		region.SetTop(geo.Y() * ratio)
		region.SetWidth(geo.Width() * ratio)
		region.SetHeight(geo.Height() * ratio)

		image := full_image.Copy(region)
		images = append(images, image)
	}

	return images, nil
}
