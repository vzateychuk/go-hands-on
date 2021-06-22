package client

type UserClient struct {
	ID   string
	Name string
}

type TxClient struct {
	ID          string
	User        *UserClient
	AccountFrom string
	AccountTo   string
	Amount      float64
}
