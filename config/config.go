package config

type Config struct {
	HttpPort string
}

func New() Config {
	return Config{
		HttpPort: ":8080",
	}
}
