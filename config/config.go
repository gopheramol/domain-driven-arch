package config

import "github.com/spf13/viper"

// Config represents the configuration for the application.
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`      // Database driver, e.g., "postgres", "mysql", etc.
	DBSource      string `mapstructure:"DB_SOURCE"`      // Database connection string or data source name.
	ServerAddress string `mapstructure:"SERVER_ADDRESS"` // Server address to bind and listen on.
}

// LoadConfig loads the application configuration from the specified path.
// The path should point to the directory where the config file is located.
func LoadConfig(path string) (config Config, err error) {
	// Add the provided path as the config search path.
	viper.AddConfigPath(path)

	// Set the name of the config file (without extension) to "app".
	viper.SetConfigName("app")

	// Set the config type to "env", indicating that it's an environment file.
	viper.SetConfigType("env")

	// Automatically check for and read environment variables with matching keys.
	viper.AutomaticEnv()

	// Read the configuration from the config file (app.env) in the provided path.
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Unmarshal the config data into the Config struct.
	err = viper.Unmarshal(&config)
	return
}
