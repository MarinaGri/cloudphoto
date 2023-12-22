package commands

import (
	"cloudphoto/internal/services"
	"fmt"
	"os"
)

func getIniConfig() services.IniConfig {
	configManager, err := services.NewConfigManager()
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

func getAwsManager(iniConfig services.IniConfig) *services.AwsManager {
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

	return awsManager
}
