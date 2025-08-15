package core

type Fetcher interface {
	Run(serviceName, kdsvcSessionId string, params any) (res []byte, err error)
}
