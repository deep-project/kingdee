package adapters

import "kingdee/client"

type LoginByValidateUser struct {
	kdsvcSessionId string
}

func (login *LoginByValidateUser) Refresh(a *client.API) error {
	return nil
}

func (login *LoginByValidateUser) KDSVCSessionId() string {
	return login.kdsvcSessionId
}
