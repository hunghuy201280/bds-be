package configs

import (
	"bds-go-auth-service/common/l"
	"github.com/spf13/viper"
)

var (
	ll = l.New()
)

type Config struct {
	Port             string `mapstructure:"PORT"`
	DBUrl            string `mapstructure:"DB_URL"`
	JWTSecret        string `mapstructure:"JWT_SECRET"`
	JWTRefreshSecret string `mapstructure:"JWT_REFRESH_SECRET"`
}

func LoadConfig() *Config {

	viper.AutomaticEnv()
	viper.SetEnvPrefix("bds_auth")
	err := viper.BindEnv("DB_URL", "db_url")
	err = viper.BindEnv("PORT", "port")
	err = viper.BindEnv("JWT_SECRET", "jwt_secret")
	err = viper.BindEnv("JWT_REFRESH_SECRET", "jwt_refresh_secret")
	if err != nil {
		ll.Fatal("Failed to read viper config", l.Error(err))
	}

	//viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	var cfg = &Config{}

	err = viper.Unmarshal(cfg)
	if err != nil {
		ll.Fatal("Failed to unmarshal config", l.Error(err))
	}

	ll.Info("Config loaded", l.Object("config", cfg))
	return cfg
}
