package main

import (
	"github.com/bobanderson-malibu/dotfree/payment-gate/cryptocurrencies/bitcoin"
	"github.com/bobanderson-malibu/dotfree/payment-gate/cryptocurrencies/bitcoincash"
	"github.com/bobanderson-malibu/dotfree/payment-gate/cryptocurrencies/ethereum"
	"github.com/bobanderson-malibu/dotfree/payment-gate/db"
)

func SyncDatabase() {
	db.Database.AutoMigrate(
		&bitcoin.BitcoinWallet{},
		&bitcoin.BitcoinWalletBalance{},

		&bitcoincash.BitcoinCashWallet{},
		&bitcoincash.BitcoinCashWalletBalance{},

		&ethereum.EthereumWallet{},
		&ethereum.EthereumWalletBalance{},
	)

	ethereum.SetupViews()
	bitcoin.SetupViews()
	bitcoincash.SetupViews()
}
