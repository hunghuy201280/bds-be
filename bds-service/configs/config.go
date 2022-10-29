package configs

import (
	"bds-service/common/l"
	"github.com/spf13/viper"
	"strings"
)

var (
	ll = l.New()
)

type Config struct {
	Port             string `yaml:"port"`
	DBUrl            string `yaml:"dbUrl"`
	JWTSecret        string `yaml:"jwtSecret"`
	JWTRefreshSecret string `yaml:"jwtRefreshSecret"`
}

func LoadConfig() *Config {
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		ll.Fatal("Failed to read viper config", l.Error(err))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	var cfg = &Config{}

	err = viper.Unmarshal(cfg)
	if err != nil {
		ll.Fatal("Failed to unmarshal config", l.Error(err))
	}

	ll.Info("Config loaded", l.Object("config", cfg))
	return cfg
}
