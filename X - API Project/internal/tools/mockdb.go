package tools

import "time"

type mockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"john": {
		AuthToken: "123-abc",
		UserName:  "john_doe",
	},
	"jane": {
		AuthToken: "987-klm",
		UserName:  "jane_doe",
	},
	"biluca": {
		AuthToken: "654-vlb",
		UserName:  "bilucax",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"john": {
		Coins:    180,
		UserName: "john_doe",
	},
	"jane": {
		Coins:    280,
		UserName: "jane_doe",
	},
	"biluca": {
		Coins:    999,
		UserName: "bilucax",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 2)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
