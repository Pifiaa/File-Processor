package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server Server
		DB     DB
	}

	Server struct {
		Host string `mapstructure:"Host"`
		Port string `mapstructure:"Port"`
	}

	DB struct {
		Host     string `mapstructure:"Host"`
		Port     string `mapstructure:"Port"`
		User     string `mapstructure:"User"`
		Password string `mapstructure:"Password"`
		Name     string `mapstructure:"Name"`
	}
)

// GetConfig carga la configuración desde un archivo y retorna un puntero a Config.
func GetConfig() (*Config, error) {
	// Configurar de viper
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	// Lee el archivo de configuración
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error al leer el archivo de configuración: %s", err)
	}

	var config Config
	// Decodifica el contenido del archivo
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("Error al cargar datos de configuración: %s", err)
	}

	return &config, nil
}
