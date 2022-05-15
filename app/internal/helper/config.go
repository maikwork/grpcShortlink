package helper

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"

	"github.com/maikwork/grpcGeteway/internal/model"
)

func ReadConfig(path string) (*model.Config, error) {
	cfg := model.Config{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Don't read file")
		return nil, err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Printf("Don't unmarshal config")
		return nil, err
	}
	return &cfg, nil
}
