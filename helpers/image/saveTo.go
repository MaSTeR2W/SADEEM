package image

import (
	"image"
	"image/jpeg"
	_ "image/png"
	"os"

	"mime/multipart"
	"slices"
	"strings"

	"github.com/nfnt/resize"
)

/* var imgsPath = path.Join(absPathToMe.Get(), "/../../public/images") + "/" */

var allowedType = []string{"png", "jpeg", "jpg"}
var strAllowedType = strings.Join(allowedType, ", ")

// absFilePath: should be absolute path with file name without extension.

func IsImage(file *multipart.FileHeader, field string) error {
	var mimeType = file.Header.Get("Content-type")
	var _type, found = strings.CutPrefix(mimeType, "image/")
	if !found {
		return &UnsupportedType{
			Field:     field,
			MessageAr: "يجب أن يكون الملف صورة وليس: (" + mimeType + ").",
			MessageEn: "File should be image not: (" + mimeType + ").",
		}
	}

	if !slices.Contains(allowedType, _type) {
		return &UnsupportedType{
			Field:     field,
			MessageAr: "امتداد الملف يجب أن يكون واحداً من: (" + strAllowedType + "), وليس: " + _type + ".",
			MessageEn: "File extenstion should be one of: (" + strAllowedType + "), not: " + _type + ".",
		}
	}
	return nil
}

func SaveAsJPEGTo(file *multipart.FileHeader, absFilePath string, width, height uint) error {
	var src, err = file.Open()

	if err != nil {
		return err
	}
	defer src.Close()

	img, _, err := image.Decode(src)

	if err != nil {
		return err
	}

	if width == 0 || height == 0 {
		var size = img.Bounds().Size()
		if height == 0 && width == 0 {
			width = uint(size.X)
			height = uint(size.Y)
		} else if width == 0 {
			width = (height * uint(size.X)) / uint(size.Y)
		} else {
			height = (width * uint(size.Y)) / uint(size.X)
		}
	}

	img = resize.Thumbnail(width, height, img, resize.Lanczos3)

	writer, err := os.Create(absFilePath)

	if err != nil {
		return err
	}
	defer writer.Close()

	return jpeg.Encode(writer, img, &jpeg.Options{Quality: 100})
}
