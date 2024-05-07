package config

import (
	

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App struct {
		Host       string `envconfig:"APP_HOST" default:"localhost"`
		Port       string `envconfig:"APP_PORT" default:"8080"`
		ApiVersion string `envconfig:"API_VERSION" default:"v0"`
		AppVersion string `envconfig:"APP_VERSION" default:"v0.0.1"`
	}
	Database struct {
		MongoUri      string `envconfig:"MONGODB_URI" default:"mongodb://localhost:27017"`
		AccessTimeout int    `envconfig:"MONGODB_ACCESS_TIMEOUT" default:"5"`
	}
}




// It is initialized once when the application starts.
var appConfig = &Config{}

// This function provides access to the appConfig variable.
func AppConfig() *Config {
	return appConfig
}

// LoadConfig loads environment variables and populates appConfig.
// It first attempts to load variables from a .env file and then
// processes environment variables according to the Config struct tags.
func LoadConfig() error {
	godotenv.Load()
	if err := envconfig.Process("", appConfig); err != nil {
		return err
	}

	

	return nil
}



