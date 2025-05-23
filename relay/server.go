package relay

import (
	"fmt"
	"net/http"

	"github.com/paywithclay/clay-relay/relay/handlers"
)

func StartServer() {
	http.HandleFunc("/charge", handlers.ChargeHandler)
	http.HandleFunc("/verify", handlers.VerifyHandler)
	http.HandleFunc("/health", handlers.HealthCheck)

	port := "3001"
	fmt.Printf("🔌 Listening on port :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
