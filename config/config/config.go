package config

type ConfigList struct {
	Token struct {
		Secret Config
	}
	Port string
}

type Config struct {
	Value string
}
