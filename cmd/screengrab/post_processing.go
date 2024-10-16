package screengrab

import (
	"textlens/internal/lib"

	"github.com/therecipe/qt/gui"
)

/*
Split full desktop image into list of images per screen.
Also resizes screens according to image:virtual-geometry ratio.
*/
func SplitFullDesktopToScreens(full_image *gui.QImage) []*gui.QImage {
	if full_image == nil {
		lib.LogError.Fatalln("Invalid image ptr provided. Exiting...")
		return nil
	}

	virtual_geometry := gui.QGuiApplication_PrimaryScreen().VirtualGeometry()
	ratio := full_image.Rect().Width() / virtual_geometry.Width()

	lib.LogInfo.Printf("Virtual geometry width: %d", virtual_geometry.Width())
	lib.LogInfo.Printf("Image width: %d", full_image.Rect().Width())
	lib.LogInfo.Printf("Resize ratio: %d", ratio)

	images := []*gui.QImage{full_image}
	// TODO: Implement support for multimonitors by capturing on all screens. QGuiApplication_Screens() throws for some reason
	// for _, screen := range gui.QGuiApplication_Screens() {
	// 	geo := screen.Geometry()
	// 	region := core.NewQRect()
	// 	region.SetLeft(geo.X() * ratio)
	// 	region.SetTop(geo.Y() * ratio)
	// 	region.SetWidth(geo.Width() * ratio)
	// 	region.SetHeight(geo.Height() * ratio)
	//
	// 	image := full_image.Copy(region)
	// 	images = append(images, image)
	// }
	return images
}
