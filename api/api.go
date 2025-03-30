package api

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code    int
	Balance int64
}

// error resopnse
type Error struct {
	Code    int
	Message int
}
