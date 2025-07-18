package adapters

import "kingdee/client"

type LoginBySign struct {
	kdsvcSessionId string
}

func (login *LoginBySign) Refresh(a *client.API) error {
	return nil
}

func (login *LoginBySign) KDSVCSessionId() string {
	return login.kdsvcSessionId
}
