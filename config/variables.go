package config

type Config struct {
    Address  string `json:"address"`
    Port     int    `json:"port"`
    Database string `json:"database"`
}

var Conf Config