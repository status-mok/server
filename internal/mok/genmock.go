//go:generate mockery --name=(.+)Mock --case=underscore
package mok

type RouteMock interface {
	Route
}

type ServerMock interface {
	Server
}

type ServerStorageMock interface {
	ServerStorage
}
