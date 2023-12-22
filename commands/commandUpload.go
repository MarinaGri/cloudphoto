package commands

import (
	"cloudphoto/internal/services"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"sync"
)

var CommandUpload = &cobra.Command{
	Use:   "upload",
	Run:   initUpload,
	Short: "Upload photos to bucket",
}

func initUpload(cmd *cobra.Command, _ []string) {
	album, _ := cmd.Flags().GetString("album")
	path, _ := cmd.Flags().GetString("path")

	photos := getPhotosFromDirectory(path)

	iniConfig := getIniConfig()

	uploadPhotos(iniConfig, photos, album)
}

func uploadPhotos(iniConfig services.IniConfig, photos []string, album string) {
	awsManager := getAwsManager(iniConfig)

	wg := sync.WaitGroup{}
	wg.Add(len(photos))
	for _, photo := range photos {
		go func(photo string) {
			defer wg.Done()
			photoKey := services.GetPhotoKey(album, filepath.Base(photo))
			data, err := os.ReadFile(photo)
			if err != nil {
				fmt.Printf("File %v can not be read\n", photo)
				return
			}

			err = awsManager.UploadPhoto(iniConfig.Bucket, photoKey, data)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("File %v successfully uploaded with key %v\n", photo, photoKey)
			}
		}(photo)
	}

	wg.Wait()
}

func getPhotosFromDirectory(path string) []string {
	var jpegFiles []string

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := filepath.Ext(file.Name())
		if ext == ".jpg" || ext == ".jpeg" {
			jpegFiles = append(jpegFiles, filepath.Join(path, file.Name()))
		}
	}

	if len(jpegFiles) == 0 {
		fmt.Printf("In directory %v there is no photos\n", path)
		os.Exit(1)
	}

	return jpegFiles
}
