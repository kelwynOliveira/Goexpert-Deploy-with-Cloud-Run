package configs

import "github.com/spf13/viper"

type Conf struct {
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
}

func LoadConfig(path string) (*Conf, error) {
	var c *Conf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	// viper.AddConfigPath(path)
	viper.SetConfigFile(path + ".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}

	return c, nil
}
