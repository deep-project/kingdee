package adapters

import "kingdee/client"

type LoginByAppSecret struct {
	kdsvcSessionId string
}

func (login *LoginByAppSecret) Refresh(a *client.API) error {
	return nil
}

func (login *LoginByAppSecret) KDSVCSessionId() string {
	return login.kdsvcSessionId
}
