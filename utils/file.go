package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

func IsImage(content []byte, typeExt []string) bool {
	s := http.DetectContentType(content)
	for _, e := range typeExt {
		if s == e {
			return true
		}
	}
	return false

}

func CropImage(url string, width, height int) {
	img, err := imaging.Open(Dir() + "/" + url)
	if err != nil {
		log.Println("Error open image => ", err)
	}
	centercropimg := imaging.Fill(img, width, height, imaging.Center, imaging.Lanczos)
	err = imaging.Save(centercropimg, Dir()+"/"+url)
	if err != nil {
		fmt.Println(err)
	}
}

func CreatePathFile(filename string) (*os.File, string, error) {
	upldir, uploadDirFull := UploadDir()
	ext := filepath.Ext(filename)
	un := fmt.Sprintf("%s%s", uuid.New(), ext)
	fn := fmt.Sprintf("%s/%s", upldir, un)
	dst, err := os.Create(fmt.Sprintf("%s/%s", uploadDirFull, un))
	if err != nil {
		return nil, "", err
	}
	return dst, fn, nil
}

func IsFileExtensionAllowed(filename string, extension []string) bool {
	ext := filepath.Ext(filename)
	for _, e := range extension {
		if ext == e {
			return true
		}
	}
	return false
}
