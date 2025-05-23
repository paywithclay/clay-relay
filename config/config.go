package config

type GatewayConfig struct {
    Gateway   string `json:"gateway"`
    SecretKey string `json:"secret_key"`
    Enabled   bool   `json:"enabled"`
}