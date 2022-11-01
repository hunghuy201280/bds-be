package configs

import (
	"bds-service/common/l"
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
	S3BucketName     string `mapstructure:"S3_BUCKET_NAME"`
	S3Region         string `mapstructure:"S3_REGION"`
	S3APIKey         string `mapstructure:"S3_API_KEY"`
	S3Secret         string `mapstructure:"S3_SECRET"`
	S3Domain         string `mapstructure:"S3_DOMAIN"`
}

func LoadConfig() *Config {

	viper.AutomaticEnv()
	viper.SetEnvPrefix("bds_api")
	err := viper.BindEnv("DB_URL", "DB_URL")
	err = viper.BindEnv("PORT", "PORT")
	err = viper.BindEnv("JWT_SECRET", "JWT_SECRET")
	err = viper.BindEnv("JWT_REFRESH_SECRET", "JWT_REFRESH_SECRET")
	err = viper.BindEnv("S3_BUCKET_NAME", "S3_BUCKET_NAME")
	err = viper.BindEnv("S3_REGION", "S3_REGION")
	err = viper.BindEnv("S3_API_KEY", "S3_API_KEY")
	err = viper.BindEnv("S3_SECRET", "S3_SECRET")
	err = viper.BindEnv("S3_DOMAIN", "S3_DOMAIN")
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
