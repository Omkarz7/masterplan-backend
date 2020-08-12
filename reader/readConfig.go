package reader

import (
	"encoding/json"
	"io/ioutil"
	"masterplan-backend/models"
	"os"

	"github.com/akshaybharambe14/go-jsonc"
)

func LoadConfiguration() (err error) {
	file, err := os.Getwd()
	if err != nil {
		return err
	}
	file = file + "/config.jsonc"
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return err
	}
	jsonDecoder := jsonc.NewDecoder(configFile)
	jsonByteStream, err := ioutil.ReadAll(jsonDecoder)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonByteStream, &models.Config)
	if err != nil {
		return err
	}
	return nil
}
