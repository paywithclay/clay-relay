package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/paywithclay/clay-relay/config"
	"github.com/paywithclay/clay-relay/relay"
)

func Setup() {
	fmt.Println("🛠️ Running Setup Mode")

	var gateway, apiKey string

	fmt.Print("Enter gateway (paystack/flutterwave): ")
	fmt.Scanln(&gateway)

	fmt.Print("Enter API secret key: ")
	fmt.Scanln(&apiKey)

	password := os.Getenv("CLAY_RELAY_PASSWORD")
	if password == "" {
		fmt.Print("Enter encryption password: ")
		fmt.Scanln(&password)
	}

	encryptedKey, err := config.Encrypt(apiKey, password)
	if err != nil {
		fmt.Printf("❌ Failed to encrypt key: %v\n", err)
		os.Exit(1)
	}

	cfg := config.GatewayConfig{
		Gateway:   gateway,
		SecretKey: encryptedKey,
		Enabled:   true,
	}

	data, _ := json.MarshalIndent(cfg, "", "  ")
	err = os.WriteFile("gateways.json.enc", data, 0644)
	if err != nil {
		fmt.Printf("❌ Failed to save config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Config saved securely.")
}

func Start() {
	fmt.Println("📡 Starting Clay Relay...")
	relay.StartServer()
}
