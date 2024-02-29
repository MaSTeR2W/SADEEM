package image

import (
	"image"
	"image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"path"
	"time"

	"mime/multipart"
	"slices"
	"strings"

	"github.com/MaSTeR2W/SADEEM/helpers/absPath"
	"github.com/MaSTeR2W/SADEEM/helpers/hprFns"
	"github.com/nfnt/resize"
)

var ImgsFolderPath = path.Join(absPath.ToMe(), "/../../public/imgs") + "/"

var allowedType = []string{"png", "jpeg", "jpg"}
var strAllowedType = strings.Join(allowedType, ", ")

// absFilePath: should be absolute path with file name without extension.

func IsImage(file *multipart.FileHeader, field string, lang string) (string, error) {
	var mimeType = file.Header.Get("Content-type")
	var _type, found = strings.CutPrefix(mimeType, "image/")
	if !found {
		return "", &UnsupportedType{
			Field:   field,
			Message: isNotImageErr(mimeType, lang),
		}
	}

	if !slices.Contains(allowedType, _type) {
		return "", &UnsupportedType{
			Field:   field,
			Message: unSupportedExtension(_type, lang),
		}
	}
	return _type, nil
}

func SaveImage(file *multipart.FileHeader, field string, width, height uint, lang string) (string, error) {

	var extension, err = IsImage(file, field, lang)

	if err != nil {
		return "", err
	}

	src, err := file.Open()

	if err != nil {
		return "", err
	}
	defer src.Close()

	img, _, err := image.Decode(src)

	if err != nil {
		return "", err
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

	rdm, err := hprFns.RdmHexStr(12)
	if err != nil {
		return "", err
	}
	var filename = time.Now().Format("20060102_150405") + "_" + rdm + "." + extension

	writer, err := os.Create(ImgsFolderPath + filename)

	if err != nil {
		return "", err
	}
	defer writer.Close()

	if extension == "png" {
		return filename, png.Encode(writer, img)
	} else {
		return filename, jpeg.Encode(writer, img, &jpeg.Options{Quality: 100})
	}
}

func isNotImageErr(mimeType string, lang string) string {
	if lang == "ar" {
		return "يجب أن يكون الملف صورة وليس: (" + mimeType + ")"
	}

	return "File should be image not: (" + mimeType + ")"
}

func unSupportedExtension(_type, lang string) string {
	if lang == "ar" {
		return "امتداد الملف يجب أن يكون واحداً من: (" + strAllowedType + "), وليس: (" + _type + ")"
	}

	return "File extenstion should be one of: (" + strAllowedType + "), not: (" + _type + ")"
}
