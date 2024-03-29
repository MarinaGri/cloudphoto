﻿package services

import (
	"cloudphoto/internal/constants"
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"os/user"
	"path/filepath"
)

type ConfigManager struct {
	user *user.User
}

type IniConfig struct {
	Bucket      string
	AccessKey   string
	SecretKey   string
	Region      string
	EndpointURL string
}

func NewConfigManager() (*ConfigManager, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	return &ConfigManager{user: currentUser}, nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (cm ConfigManager) GetConfigFilePath() string {
	return filepath.Join(cm.user.HomeDir, ".config", "cloudphoto", "cloudphotorc", "cloudphoto.ini")
}

func (cm ConfigManager) GenerateIni(config IniConfig) error {
	path := cm.GetConfigFilePath()

	cfg := ini.Empty()

	section, err := cfg.NewSection(constants.DefaultSectionName)
	if err != nil {
		return err
	}

	section.NewKey(constants.Bucket, config.Bucket)
	section.NewKey(constants.AccessKey, config.AccessKey)
	section.NewKey(constants.SecretKey, config.SecretKey)
	section.NewKey(constants.Region, config.Region)
	section.NewKey(constants.EndpointURL, config.EndpointURL)

	if !fileExists(path) {
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}

		file, err := os.Create(path)
		_ = file.Close()
		if err != nil {
			return err
		}
	}

	err = cfg.SaveTo(path)
	if err != nil {
		return err
	}

	fmt.Printf("Ini file successfully created at %v\n", path)

	return nil
}

func (cm ConfigManager) TryGetConfig() (bool, *IniConfig, error) {
	path := cm.GetConfigFilePath()

	cfg, err := ini.Load(path)
	if err != nil {
		return false, nil, err
	}

	section := cfg.Section(constants.DefaultSectionName)
	config := &IniConfig{
		Bucket:      section.Key(constants.Bucket).String(),
		AccessKey:   section.Key(constants.AccessKey).String(),
		SecretKey:   section.Key(constants.SecretKey).String(),
		Region:      section.Key(constants.Region).String(),
		EndpointURL: section.Key(constants.EndpointURL).String(),
	}

	return cm.isValidConfig(*config), config, nil
}

func (cm ConfigManager) isValidConfig(config IniConfig) bool {
	return len(config.Bucket) > 0 &&
		len(config.AccessKey) > 0 &&
		len(config.SecretKey) > 0 &&
		len(config.Region) > 0 &&
		len(config.EndpointURL) > 0
}
