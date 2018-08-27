package tasks

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// ImageResizetask ...
type ImageResizetask struct {
	BaseImageName      string
	ImageFileExtension string
}

// NewImageResizeTask ...
func NewImageResizeTask(baseImageName string, imageFileExtension string) *ImageResizetask {
	return &ImageResizetask{BaseImageName: baseImageName, ImageFileExtension: imageFileExtension}
}

// Perform ...
func (t *ImageResizetask) Perform() {
	thumbImageFilePath := t.BaseImageName + "_thumb.png"
	fmt.Println("Creating new thumbnail at ", thumbImageFilePath)

	originalimagefile, err := os.Open(t.BaseImageName + t.ImageFileExtension)

	if err != nil {
		log.Println(err)
		return
	}

	img, err := png.Decode(originalimagefile)

	if err != nil {
		log.Println("Encountered Error while decoding image file: ", err)
	}

	thumbImage := resize.Resize(270, 0, img, resize.Lanczos3)
	thumbImageFile, err := os.Create(thumbImageFilePath)

	if err != nil {
		log.Println("Encountered error while resizing image:", err)
	}

	png.Encode(thumbImageFile, thumbImage)
}
