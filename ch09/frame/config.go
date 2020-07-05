package frame

import "github.com/spf13/viper"

func LoadConfig(path, configName, configType string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
