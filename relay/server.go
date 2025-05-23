package relay

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/charge", ChargeHandler)
	http.HandleFunc("/verify", VerifyHandler)
	http.HandleFunc("/health", HealthCheck)

	port := "3001"
	fmt.Printf("🔌 Listening on port :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
