//go:generate mockery --name=(.+)Mock --case=underscore
package mok

type EndpointMock interface {
	Endpoint
}

type ServerMock interface {
	Server
}

type ServerStorageMock interface {
	ServerStorage
}
