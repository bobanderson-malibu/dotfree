package bitcoin

import (
	"time"

	"github.com/bobanderson-malibu/dotfree/payment-gate/settings"
)

func TaskUpdateBitcoinWalletBalances() {
	c := time.Tick(24 * time.Hour)
	for range c {
		UpdateBitcoinWalletBalances()
	}
}

func TaskUpdateBitcoinTransactionFee() {
	UpdateBTCTransactionFee()
	c := time.Tick(1 * time.Minute)
	for range c {
		UpdateBTCTransactionFee()
	}
}

func init() {
	if !settings.APPLICATION_SETTINGS.Debug {
		go TaskUpdateBitcoinWalletBalances()
		go TaskUpdateBitcoinTransactionFee()
	}
}
