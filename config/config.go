package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	RestServer RestServer `mapstructure:"rest_server"`
	API        API        `mapstructure:"api"`
	OAuth      OAuth      `mapstructure:"oauth"`
}

type RestServer struct {
	Name            string `mapstructure:"name"`
	Addr            string `mapstructure:"addr"`
	ShutdownTimeout int    `mapstructure:"shutdown_timeout"` // giây
}

type API struct {
	Version string `mapstructure:"version"`
	Env     string `mapstructure:"env"`
}

type OAuth struct {
	Issuer   string `mapstructure:"issuer"`
	Audience string `mapstructure:"audience"`
	KeyID    string `mapstructure:"key_id"`
	TokenTTL int    `mapstructure:"token_ttl"` // giây
}

// NewConfig đọc config/config.yml, biến môi trường ghi đè theo quy tắc "." -> "_"
// (ví dụ REST_SERVER_ADDR=:9001, OAUTH_TOKEN_TTL=60).
// File được tìm ở ./config, nên chạy từ thư mục service (go run ./cmd/rest) hoặc trong Docker.
func NewConfig() (*Config, error) {
	var cfg Config

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("config")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
