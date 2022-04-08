package config

type ConfigList struct {
	Token struct {
		Secret Config
	}
}

type Config struct {
	Value string
}
