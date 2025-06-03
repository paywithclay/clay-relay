# 🧱 Clay Relay

> *Your private tunnel to every payment gateway.*  
> 🔐 Built for **PayWithClay** — no keys, no trust, just payments.

---

## 🤝 What is Clay Relay?
> This is still in experimental development please

**Clay Relay** is a self-hosted, secure proxy service that lets you connect any payment gateway (like Paystack, Flutterwave, or Stripe) without sharing your API keys with anyone — not even us.

It acts as your own personal bridge between the **Clay SDK** and your favorite payment providers.

> Think of it like a **WalletConnect**, but for traditional finance. 💸

---

## 🎯 Why Use Clay Relay?

- ✅ **Zero Key Exposure**: You keep your secrets — we never see them.
- 🔐 **Encrypted Configuration**: Your gateway keys are AES-256 encrypted at rest.
- 🛡️ **Secure Requests**: HMAC-signed requests ensure only authorized calls reach your relay.
- 🌍 **Multi-Gateway Ready**: Connect Paystack, Flutterwave, Stripe — all from one interface.
- 🚀 **Easy Setup**: Run it locally, on a VPS, or in Docker — it's up to you.

---

## 📦 Features

- AES-256 encryption for config files
- HMAC request verification for secure communication
- Self-contained microservice architecture
- Easy CLI setup and start commands
- Lightweight Go binary – fast and efficient
- Ready for integration with the **Clay SDK**

---

## 🧪 Getting Started

### 1. Clone the repo (if needed)

```bash
git clone https://github.com/paywithclay/clay-relay.git
cd clay-relay
```

> ⚠️ If you're not using GitHub, skip this and work in your local folder.

---

### 2. Initialize Go Module

```bash
go mod init github.com/paywithclay/clay-relay
go mod tidy
```

---

### 3. Build the Binary

```bash
go build -o clay-relay
```

---

### 4. Set Up Your Gateways

```bash
CLAY_RELAY_SECRET=mysecret CLAY_RELAY_PASSWORD=mypassword ./clay-relay setup
```

You'll be asked:
- Which gateway you want to connect (e.g., `paystack`)
- Your secret key for that gateway
- An encryption password (or use the env var)

This creates a secure, encrypted config file: `gateways.json.enc`

---

### 5. Start the Relay

```bash
CLAY_RELAY_SECRET=mysecret ./clay-relay start
```

✅ Now your relay is running on `http://localhost:3001`

---

## 🛡️ Security Overview

| Feature | How It Works |
|--------|--------------|
| 🔒 Encrypted Config | Gateway keys stored using AES-256-GCM |
| 🧾 HMAC Verification | All incoming requests must be signed by the Clay SDK |
| ⏱ Replay Protection | Timestamp checks prevent replay attacks |
| 🙅‍♂️ Zero Trust | We never see or store your API keys |

---

## 🧰 Supported Endpoints

| Endpoint | Description |
|---------|-------------|
| `POST /charge` | Initiate a payment |
| `POST /verify` | Verify a transaction |
| `GET /health` | Health check |

All endpoints require valid HMAC signatures from the **Clay SDK**.

---

## 🧬 Integration with Clay SDK

Once you're running `clay-relay`, point the **Clay SDK** to your local endpoint:

```js
const clay = new Clay({
  publicKey: 'your-public-key',
  proxyUrl: 'http://localhost:3001'
});
```

Now all payment traffic goes securely through your own relay.

---

## 🐳 Optional: Run with Docker

```bash
docker build -t paywithclay/clay-relay .
docker run -p 3001:3001 \
  -e CLAY_RELAY_SECRET=mysecret \
  -v $(pwd)/gateways.json.enc:/app/gateways.json.enc \
  paywithclay/clay-relay
```

---

## 📁 File Structure

```
clay-relay/
├── main.go
├── cli/              # CLI setup commands
├── config/           # Secure config handling + encryption
├── relay/            # HTTP server + handlers
├── utils/            # HMAC signing & verification
├── adapters/         # Payment gateway plugins (coming soon)
└── gateways.json.enc # Encrypted config file
```

---

## 🧑‍💻 Want to Contribute?

We’re building the future of universal payment infrastructure — and **you're welcome to join**.

Feel free to open issues, submit PRs, or suggest new features!

---

## 📣 Questions? Feedback?

Reach out to us anytime at [support@paywithclay.co] or join our Discord community 👇

> 💬 **Join the conversation:** Coming soon🥲

---

## ❤️ Proudly Made for PayWithClay

> Because paying should feel simple — not scary.

---

🎉 You made it! Now go run your relay and start accepting payments securely with **Clay Relay**.

---
