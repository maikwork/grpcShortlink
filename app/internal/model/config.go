package model

type DB struct {
	Name string `yaml:"name"`
	URL string `yaml:"jdbc"`
}

type REST struct {
	Port string `yaml:"rest_port"`
}

type GRPC struct {
	Port string `yaml:"grpc_port"`
}
type SettingServ struct {
	GRPC GRPC `yaml:"grpc"`
	REST REST `yaml:"rest"`
}

type Config struct {
	DB          DB          `yaml:"db"`
	SettingServ SettingServ `yaml:"setting_serv"`
}
