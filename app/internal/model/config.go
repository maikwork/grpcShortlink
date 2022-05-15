package model

type DB struct {
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Pass     string `yaml:"pass"`
	DB       string `yaml:"db"`
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
