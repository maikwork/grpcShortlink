package model

type DB struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Pass     string `json:"pass"`
	DB       string `json:"db"`
}

type Config struct {
	DB DB `json:"db"`
}
