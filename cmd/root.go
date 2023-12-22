package cmd

import (
	"cloudphoto/commands"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{Use: "cloudphoto"}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	rootCmd.AddCommand(commands.CommandInit)

	rootCmd.AddCommand(commands.CommandUpload)
	commands.CommandUpload.Flags().String("album", "", "Album name")
	commands.CommandUpload.MarkFlagRequired("album")
	commands.CommandUpload.Flags().String("path", currentDir, "Path to directory with photos")

	rootCmd.AddCommand(commands.CommandDownload)
	commands.CommandDownload.Flags().String("album", "", "Album name")
	commands.CommandDownload.MarkFlagRequired("album")
	commands.CommandDownload.Flags().String("path", currentDir, "Path to directory with photos")

	rootCmd.AddCommand(commands.CommandList)
	commands.CommandList.Flags().String("album", "", "Album name")

	rootCmd.AddCommand(commands.CommandDelete)
	commands.CommandDelete.Flags().String("album", "", "Album name")
	commands.CommandDelete.MarkFlagRequired("album")
	commands.CommandDelete.Flags().String("photo", "", "Photo name to delete")

	rootCmd.AddCommand(commands.CommandMksite)
}
