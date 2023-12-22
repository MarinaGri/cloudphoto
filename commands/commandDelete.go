package commands

import (
	"cloudphoto/internal/services"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var CommandDelete = &cobra.Command{
	Use:   "delete",
	Run:   initDelete,
	Short: "Delete albums",
}

func initDelete(cmd *cobra.Command, _ []string) {
	album, _ := cmd.Flags().GetString("album")
	photo, _ := cmd.Flags().GetString("photo")

	iniConfig := getIniConfig()

	awsManager := getAwsManager(iniConfig)

	if len(photo) == 0 {
		err := awsManager.DeletePhotosByPrefix(iniConfig.Bucket, album)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("Album %v deleted", album)
	} else {
		err := awsManager.DeletePhoto(iniConfig.Bucket, services.GetPhotoKey(album, filepath.Base(photo)))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("Photo %v deleted", photo)
	}
}
