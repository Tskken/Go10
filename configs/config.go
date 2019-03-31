package configs

import (
	"encoding/json"
	"github.com/Tskken/Go10/errors"
	"os"
	"path/filepath"
)

const (
	ConfigPath = "../Go10/configs/config.json"
)

type Color struct {
	R uint16
	G uint16
	B uint16
	A uint16
}

type CardColor struct {
	Name  string
	Color Color
}

type CardType struct {
	Name     string
	HasColor bool
	Score    uint8
	Count    uint8
}

type Config struct {
	ColorData []CardColor
	CardData  []CardType
}

func LoadConfig() (*Config, error) {
	path, err := filepath.Abs(ConfigPath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer errors.FileClose(file)

	dec := json.NewDecoder(file)

	config := new(Config)

	if err = dec.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
