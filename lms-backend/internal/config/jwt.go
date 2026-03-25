package config

import (
	"time"
)

// JWTConfig holds JWT configuration
type JWTConfig struct {
	SecretKey        string
	RefreshSecretKey string
	AccessExpiry     time.Duration
	RefreshExpiry    time.Duration
}

// LoadJWTConfig loads JWT configuration from environment
func LoadJWTConfig() *JWTConfig {
	return &JWTConfig{
		SecretKey:        MustGetEnv("JWT_SECRET_KEY"),
		RefreshSecretKey: MustGetEnv("JWT_REFRESH_SECRET"),
		AccessExpiry:     GetEnvDuration("JWT_ACCESS_EXPIRY", 15*time.Minute),
		RefreshExpiry:    GetEnvDuration("JWT_REFRESH_EXPIRY", 7*24*time.Hour),
	}
}
