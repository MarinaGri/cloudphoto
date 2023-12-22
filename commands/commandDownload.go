package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CommandDownload = &cobra.Command{
	Use:   "download",
	Run:   initDownload,
	Short: "Download photos from bucket",
}

func initDownload(cmd *cobra.Command, _ []string) {
	album, _ := cmd.Flags().GetString("album")
	path, _ := cmd.Flags().GetString("path")

	iniConfig := getIniConfig()

	awsManager := getAwsManager(iniConfig)

	err := awsManager.DownloadPhotos(iniConfig.Bucket, album, path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
