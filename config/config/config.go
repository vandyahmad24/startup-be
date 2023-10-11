package config

type ConfigList struct {
	Token struct {
		Secret Config
	}
	Port              string
	MidtransServerKey string
	MidtransClientKey string
}

type Config struct {
	Value string
}
