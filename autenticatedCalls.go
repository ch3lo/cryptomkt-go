package client

// Account gives the account information of the client.
// https://developers.cryptomkt.com/es/#cuenta
func (client *Client) Account() (string, error) {
	return client.get("account", nil)
}

// Balance returns the actual balance of the wallets of the client in cryptomarket
// https://developers.cryptomkt.com/es/#obtener-balance
func (client *Client) Balance() (string, error) {
	return client.get("balance", nil)
}

// Wallets is an alias for Balance
// https://developers.cryptomkt.com/es/#obtener-balance
func (client *Client) Wallets() (string, error) {
	return client.Balance()
}

func (client *Client) Transactions(args map[string]interface{}) (string, error) {
	return client.get("transactions", args)
}

// ActiveOrders returns the active orders of the client
// given a market and an opitonal page number
// https://developers.cryptomkt.com/es/#ordenes-activas
func (client *Client) ActiveOrders(args map[string]interface{}) (string, error) {
	return client.get("orders/active", args)
}

// ExecutedOrders return a list of the executed orders of the client
// given a market and an optional page
// https://developers.cryptomkt.com/es/#ordenes-ejecutadas
func (client *Client) ExecutedOrders(args map[string]interface{}) (string, error) {
	return client.get("orders/executed", args)
}


// OrderStatus gives the status of an order
// given the order "id"
// https://developers.cryptomkt.com/es/#estado-de-orden
func (client *Client) OrderStatus(args map[string]interface{}) (string, error) {
	return client.get("orders/status", args)
}

// Instant emulates an order in the current state of the Instant Exchange of CryptoMarket
// given a "market", a "type" ("buy" or "sell") and an "amount"
// https://developers.cryptomkt.com/es/#obtener-cantidad
func (client *Client) Instant(args map[string]interface{}) (string, error) {
	return client.get("orders/instant/get", args)
}

// CreateOrder signal to create an order of buy or sell in CryptoMarket
// given an "amount", a "market", a "price" and a "type" ("buy" or "sell")
// https://developers.cryptomkt.com/es/#crear-orden
func (client *Client) CreateOrder(args map[string]interface{}) (string, error) {
	return client.post("orders/create", args)
}

// CancelOrder signal to cancel an order in CryptoMarket
// given the order "id"
// https://developers.cryptomkt.com/es/#cancelar-una-orden
func (client *Client) CancelOrder(args map[string]interface{}) (string, error) {
	return client.post("orders/cancel", args)
}

// CreateInstant makes an order in the Instant Exchange of CryptoMarket
// given a "market", a "type" ("buy" or "sell") and an "amount"
// https://developers.cryptomkt.com/es/#crear-orden-2
func (client *Client) CreateInstant(args map[string]interface{}) (string, error) {
	return client.post("orders/instant/create", args)
}

// RequestDeposit notifies a deposit to a wallet of local currency
// given an "amount", a "bank_account", a "date",
// a "tracking_code" (only for mexico)
// and a "voucher" (for Mexico, Brasil and European Union only)
// https://developers.cryptomkt.com/es/#notificar-deposito
func (client *Client) RequestDeposit(args map[string]interface{}) (string, error) {
	return client.post("request/deposit", args)
}

// Request withdrawal notifies a withdrawal from a bank account of the client
// given a "bank_account" and an "amount"
// https://developers.cryptomkt.com/es/#notificar-retiro
func (client *Client) RequestWithdrawal(args map[string]interface{}) (string, error) {
	return client.post("request/withdrawal", args)
}

// Transfer move crypto between wallets
// given the "adress" of the destiny wallet,
// the "amount", a "currency" of the origin wallet,
// and an optional "memo"
// https://developers.cryptomkt.com/es/#transferir
func (client *Client) Transfer(args map[string]interface{}) (string, error) {
	return client.post("transfer", args)
}

// NewOrder enables a payment order, and gives a QR and urls
// https://developers.cryptomkt.com/es/#crear-orden-de-pago
func (client *Client) NewOrder(args map[string]interface{}) (string, error) {
	return client.post("payment/new_order", args)
}

// CreateWallet creates a wallet to pay a payment order
// https://developers.cryptomkt.com/es/#crear-billetera-de-orden-de-pago
func (client *Client) CreateWallet(args map[string]interface{}) (string, error) {
	return client.post("payment/create_wallet", args)
}

// PaymentOrders returns all the generated payment orders
// https://developers.cryptomkt.com/es/#listado-de-ordenes-de-pago
func (client *Client) PaymentOrders(args map[string]interface{}) (string, error) {
	return client.get("payment/orders", args)
}

// PaymentStatus gives the status of a pyment order
// given the order "id"
// https://developers.cryptomkt.com/es/#estado-de-orden-de-pago
func (client *Client) PaymentStatus(args map[string]interface{}) (string, error) {
	return client.get("payment/status", args)
}
