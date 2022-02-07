package utils

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

const Photo_Path = "dir/photo/"
const File_Path = "dir/file/"

// "file_*.csv"
// dir/photo/{date}/img_{user_id}_*.jpg
func MoveTmpFile(file *multipart.FileHeader, path string, formatName string) (*os.File, error) {

	if strings.Contains(path, Photo_Path) {
		createDateFolder(path)
	}

	tempFile, err := ioutil.TempFile(path, formatName)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer tempFile.Close()
	// write this byte array to our temporary file

	dataFile, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	byteContainer, err := ioutil.ReadAll(dataFile)
	tempFile.Write(byteContainer)
	return tempFile, nil
}

func GetFormatPhotoName(userId string) string {
	return "img_" + userId + ".jpg"
}

func GetFormatFileName(fileExt string) string {
	return "file_" + "*." + fileExt
}

func GetDate() string {
	date := time.Now().Local()
	return date.Format("01-10-2006")
}

func createDateFolder(path string) {
	if exists, _ := exists(path); !exists {
		os.MkdirAll(path, 0777)
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
