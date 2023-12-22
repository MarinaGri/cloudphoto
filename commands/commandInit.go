package commands

import (
	"cloudphoto/internal/constants"
	"cloudphoto/internal/services"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CommandInit = &cobra.Command{
	Use:   "init",
	Run:   initFunc,
	Short: "Initialize command",
}

func initFunc(_ *cobra.Command, _ []string) {
	bucket := scanValue("Enter bucket")
	accessKey := scanValue("Enter access key id")
	secretKey := scanValue("Enter secret access key")
	iniConfig := services.IniConfig{
		Bucket:      bucket,
		AccessKey:   accessKey,
		SecretKey:   secretKey,
		Region:      constants.CurrentRegion,
		EndpointURL: constants.CurrentEndpointURL,
	}

	iniConfigFromFile := generateOrUpdateIni(iniConfig)

	createBucketIfNotExist(iniConfigFromFile)
}

func scanValue(printingValue string) string {
	var result string
	fmt.Println(printingValue)
	_, err := fmt.Scan(&result)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return result
}

func generateOrUpdateIni(iniConfig services.IniConfig) services.IniConfig {
	configManager, err := services.NewConfigManager()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = configManager.GenerateIni(iniConfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	isValid, iniConfigFromFile, err := configManager.TryGetConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if !isValid {
		fmt.Println("Ini config file is not valid")
		os.Exit(1)
	}

	return *iniConfigFromFile
}

func createBucketIfNotExist(iniConfig services.IniConfig) {
	awsConfig := services.AwsConfig{
		AccessKey:   iniConfig.AccessKey,
		SecretKey:   iniConfig.SecretKey,
		Region:      iniConfig.Region,
		EndpointURL: iniConfig.EndpointURL,
	}

	awsManager, err := services.NewAwsManager(awsConfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	exists, err := awsManager.BucketExists(iniConfig.Bucket)
	if err != nil {
		fmt.Printf("Bucket with name %v already exists\n", iniConfig.Bucket)
		os.Exit(1)
	}

	if !exists {
		err := awsManager.CreateBucket(iniConfig.Bucket)
		if err != nil {
			fmt.Printf("Can not create bucket with name %v\n", iniConfig.Bucket)
			os.Exit(1)
		}

		fmt.Printf("Bucket with name '%v' created\n", iniConfig.Bucket)
	} else {
		fmt.Printf("Bucket with name '%v' exists\n", iniConfig.Bucket)
	}
}
