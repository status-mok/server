//go:generate mockery --name=(.+)Mock --case=underscore
package mok

type ServerMock interface {
	Server
}

type ServerStorageMock interface {
	ServerStorage
}
