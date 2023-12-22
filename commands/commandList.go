package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CommandList = &cobra.Command{
	Use:   "list",
	Run:   initList,
	Short: "View the list of albums and photos in the album",
}

func initList(cmd *cobra.Command, _ []string) {
	album, _ := cmd.Flags().GetString("album")

	iniConfig := getIniConfig()

	awsManager := getAwsManager(iniConfig)

	var result []string
	var err error
	if len(album) == 0 {
		result, err = awsManager.GetPrefixes(iniConfig.Bucket)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	} else {
		result, err = awsManager.GetPhotos(iniConfig.Bucket, album)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	for _, val := range result {
		fmt.Println(val)
	}
}
