package bitcoincash

import (
	"time"

	"github.com/bobanderson-malibu/dotfree/payment-gate/settings"
)

func TaskUpdateBitcoinCashWalletBalances() {
	c := time.Tick(24 * time.Hour)
	for range c {
		UpdateBitcoinCashWalletBalances()
	}
}

func TaskUpdateBitcoinCashTransactionFee() {
	UpdateBCHTransactionFee()
	c := time.Tick(1 * time.Minute)
	for range c {
		UpdateBCHTransactionFee()
	}
}

func init() {
	if !settings.APPLICATION_SETTINGS.Debug {
		go TaskUpdateBitcoinCashTransactionFee()
	}
}
