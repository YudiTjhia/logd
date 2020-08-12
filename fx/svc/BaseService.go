package svc

const (
	ACCOUNT_ID_IS_INVALID = "account_id_is_invalid"
	LOGIN_ID_IS_INVALID = "login_id_is_invalid"
)
type BaseService struct {
}

func(service BaseService) ValidtAccID(accountID string) error {
	return StringRequired("account_id", accountID)
}

func(service BaseService) ValidtLoginID(loginID string) error {
	return StringRequired("login_id", loginID)
}


