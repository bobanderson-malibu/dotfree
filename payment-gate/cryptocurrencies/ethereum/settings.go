package ethereum

import (
	"github.com/bobanderson-malibu/dotfree/payment-gate/settings"
)

var (
	PAYMENT_GATE_SETTINGS = settings.LoadPaymentGateSettings()
	WEI_IN_ETH            = float64(1e18)
)
