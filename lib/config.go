/*
 * This file is part of the dupman/server project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>, February 2022
 */

package lib

import (
	"time"

	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

// ServerConfig represents Server config.
type ServerConfig struct {
	ListenAddr     string   `mapstructure:"SERVER_ADDR" default:"0.0.0.0"`
	Port           int      `mapstructure:"SERVER_PORT" default:"8080"`
	TrustedProxies []string `mapstructure:"TRUSTED_PROXIES"`
}

// DatabaseConfig represents Database config.
type DatabaseConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Database string `mapstructure:"DB_NAME"`
	Port     int    `mapstructure:"DB_PORT"`
}

// JWTConfig represents JWT config.
type JWTConfig struct {
	PrivateKeyLocation string        `mapstructure:"JWT_PRIVATE_KEY_LOCATION"`
	PublicKeyLocation  string        `mapstructure:"JWT_PUBLIC_KEY_LOCATION"`
	Expiration         time.Duration `mapstructure:"JWT_EXPIRATION" default:"48h"`
}

// CORSConfig represents CORS config.
type CORSConfig struct {
	Origins []string `mapstructure:"CORS_ORIGINS" default:"*"`
	Methods []string `mapstructure:"CORS_METHODS" default:"GET,POST,PUT,HEAD,OPTIONS,DELETE"`
	Headers []string `mapstructure:"CORS_HEADERS" default:"*"`
}

// Config represents General config.
type Config struct {
	Env      string         `mapstructure:"ENV" default:"prod"`
	Server   ServerConfig   `mapstructure:",squash"`
	Database DatabaseConfig `mapstructure:",squash"`
	JWT      JWTConfig      `mapstructure:",squash"`
	CORS     CORSConfig     `mapstructure:",squash"`
}

// NewConfig creates a new Config.
func NewConfig() (conf Config) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	defaults.SetDefaults(&conf)

	_ = viper.ReadInConfig()
	_ = viper.Unmarshal(&conf)

	return conf
}
