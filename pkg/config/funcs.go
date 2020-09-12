package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Vstural/iconfer"
)

func DefaultConfigName() string {
	return ".mailer_conf"
}
func DefaultConfigPath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get home dir fail: %s", err.Error())
	}
	return fmt.Sprintf("%s/%s", homedir, DefaultConfigName()), nil
}

func ReadConfig(path string) (*Conf, error) {
	f, err := iconfer.ReadFileContent(path)
	if err != nil {
		return nil, err
	}
	c := &Conf{}
	err = json.Unmarshal([]byte(f), c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func LoadConfig() (*Conf, error) {
	path, err := DefaultConfigPath()
	if err != nil {
		return nil, err
	}
	return ReadConfig(path)
}
