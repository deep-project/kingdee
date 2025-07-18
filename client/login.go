package client

type Login interface {
	Refresh(*API) error
	KDSVCSessionId() string
}
