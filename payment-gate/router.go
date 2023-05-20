package main

import (
	"github.com/gocraft/web"

	"github.com/bobanderson-malibu/dotfree/payment-gate/cryptocurrencies"
	"github.com/bobanderson-malibu/dotfree/payment-gate/cryptocurrencies/bitcoin"
	"github.com/bobanderson-malibu/dotfree/payment-gate/cryptocurrencies/bitcoincash"
	"github.com/bobanderson-malibu/dotfree/payment-gate/cryptocurrencies/ethereum"

	"github.com/bobanderson-malibu/dotfree/payment-gate/exchange"
)

func ConfigureRouter(router *web.Router) *web.Router {

	bitcoinRouter := router.Subrouter(Context{}, "/bitcoin")
	bitcoin.ConfigureRouter(bitcoinRouter)

	bitcoinCashRouter := router.Subrouter(Context{}, "/bitcoin_cash")
	bitcoincash.ConfigureRouter(bitcoinCashRouter)

	ethereumRouter := router.Subrouter(Context{}, "/ethereum")
	ethereum.ConfigureRouter(ethereumRouter)

	// Exchange
	exchangeRouter := router.Subrouter(Context{}, "/exchange")
	exchange.ConfigureRouter(exchangeRouter)

	// Currency
	router.Get("/currency/:base_currency", cryptocurrencies.ViewShowCurrency)
	return router
}
