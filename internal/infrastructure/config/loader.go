package config

func Load() *Config {
	return &Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "file-uploader"),
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "file_uploader"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Storage: StorageConfig{
			BasePath: getEnv("STORAGE_BASE_PATH", "./data"),
		},
	}
}
