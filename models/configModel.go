package models

import "time"

var Config = Configuration{}

type Configuration struct {
	Port               string        `json:"Port"`
	JWTKey             string        `json:"JWTKey"`
	TokenExpiryTime    time.Duration `json:"TokenExpiryTime"`
	LogPath            string        `json:"LogPath"`
	MasterplanFilename string        `json:"MasterplanFilename"`
	Database           struct {
		DatabaseName string `json:"DatabaseName"`
		Host         string `json:"Host"`
		Username     string `json:"Username"`
		Password     string `json:"Password"`
		Flags        string `json:"Flags"`
	} `json:"Database"`
}
