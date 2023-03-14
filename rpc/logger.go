package rpc

type Logger interface {
	Errorf(template string, args ...interface{})
	Infof(template string, args ...interface{})
}
