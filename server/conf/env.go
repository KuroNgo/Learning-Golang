package conf

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	FrontEndOrigin string `mapstructure:"FRONT_END_ORIGIN"`

	JWTTokenSecret string        `mapstructure:"JWT_SECRET"`
	TokenExpireIn  time.Duration `mapstructure:"TOKEN_EXPIRE_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`

	GoogleClientID         string `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
	GoogleClientSecret     string `mapstructure:"GOOGLE_OAUTH_CLIENT_SECRET"`
	GoogleOAuthRedirectURL string `mapstructure:"GOOGLE_OAUTH_REDIRECT_URL"`

	CGO_ENABLED string `mapstructure:"CGO_ENABLED"`
}

func LoadConfig(path string) (config Config, err error) {
	// viper dùng để đọc file config và set các giá trị vào struct
	// viper cũng có thể đọc các biến môi trường
	// viper cũng có thể đọc các flag được set khi chạy ứng dụng
	// viper cũng có thể đọc các giá trị được set trong code
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
