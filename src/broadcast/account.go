package broadcast

type Account struct {
	Id        string `json:"id,omitempty"`
	AccountId string `json:"accountId,omitempty"`
}

func NewAccount() *Account {
	return &Account{}
}
