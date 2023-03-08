package config

var defaultConfig = map[string]interface{}{
	"server.port":             9090,
	"server.timeoutSecs":      20,
	"server.readTimeoutSecs":  20,
	"server.writeTimeoutSecs": 40,

	"jwt.secret":      "secret-key",
	"jwt.sessionTime": 864000,

	// "db.dataSourceName": "postgres://postgres:password@nft-market.cqtyaecn0tih.us-east-2.rds.amazonaws.com:5432/postgres?sslmode=disable",
	"db.dataSourceName":   "postgres://common:password@localhost:5432/rena?sslmode=disable",
	"db.migrate.enable":   false,
	"db.migrate.dir":      "/migrations",
	"db.pool.maxOpen":     50,
	"db.pool.maxIdle":     5,
	"db.pool.maxLifetime": 5,

	"crypto.platformFee": 100000000000000000,
	"crypto.network":     "https://rinkeby.infura.io/v3/8c9a3de5eb604d56b8ea15f081d5bdd9",
	"crypto.webSocket":   "wss://rinkeby.infura.io/ws/v3/8c9a3de5eb604d56b8ea15f081d5bdd9",
	// "crypto.network":           "http://localhost:8545",
	// "crypto.webSocket":         "ws://localhost:8545",
	"crypto.marketplace":       "0xA5aE3F3116DDf81B170eaD4CF6D60F5c6063E23E",
	"crypto.bundleMarketplace": "0x9E1fceef3e0BC0c75DBb8C2FAb9823F843594166",
	"crypto.addressRegistry":   "0x158885d75eF4cc50Da8ef1b86238dAF12c99dD62",
	"crypto.tokenRegistry":     "0xA60e35691A9fE372673C39F88cd04a3709097f5b",
	"crypto.privateKey":        "aed7147564caaea6dd4ef501c8b880c6b2c39cb5034b9c80c9eb11314449f658",
}
