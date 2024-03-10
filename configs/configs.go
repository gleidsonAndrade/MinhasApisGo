package configs

import (
	"github.com/spf13/viper"
)

var cfg *Config


type Config struct {
	DBDriver      string 'mapsturct:"DB_DRIVER"'
	DBHost        string 'mapstruct:"DB_HOST"'
	DBport        string 'mapstruct:"DB_PORT"'
	DBUser        string 'mapsruct:"DB_USER"'
	DBPassword    string 'mapstruct:"DB_PASSWORD"
	DBName        string 'mapstruct:"DB_NAME"'
	WebServerPort string 'mapstruct:"WEB_SERVER_PORT"'
	JWTSecret     string 'mapstruct:"JWT_SECRET"'
	JWTExpiresIn  int 'mapstruct:"JWT_EXPIRESIN'
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	viper.setConfigName("app_config")
	viper.setConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	error := viper.ReadInConfig()
	if error != nil {
		return nil, error
	}
	error = viper.Unmarshal(&cfg)
	if error != nil {
		panic(error)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, error
}
