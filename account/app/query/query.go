package query

type (
	// Account account document query result
	Account struct {
		ID   string
		Name string
	}

	// Accounts account collection query result
	Accounts []struct {
		Name string
	}

	// AccountQuery account query from data
	AccountQuery interface {
		FindByID(id string) Account
		FindByName(name string) Accounts
	}
)
