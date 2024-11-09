package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    JWTSecret  string
}

func LoadConfig() (Config, error) {
    var config Config
    viper.AddConfigPath(".")
    viper.SetConfigName("app")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
        return config, err
    }

    err = viper.Unmarshal(&config)
    return config, err
}
