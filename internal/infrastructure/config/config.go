package config

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Storage  StorageConfig
}

type AppConfig struct {
	Name string
	Env  string
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type StorageConfig struct {
	BasePath string
}
