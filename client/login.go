package client

type Login interface {
	KDSVCSessionId(*API) (string, error)
}
