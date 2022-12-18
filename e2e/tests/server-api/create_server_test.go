package server_api_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/status-mok/server/e2e/server"
	serverAPI "github.com/status-mok/server/pkg/server-api"
)

var ctx = context.Background()

var _ = Describe("Create server", Ordered, func() {
	var srv *server.TestServer

	BeforeAll(func() {
		srv = server.NewServer()
	})

	AfterAll(func() {
		srv.Close()
	})

	It("should finish successfully", func() {
		resp, err := srv.ServerGRPCClient().Create(ctx, &serverAPI.CreateRequest{
			Name: "123",
			Type: serverAPI.ServerType_SERVER_TYPE_HTTP,
		})

		Expect(err).To(BeNil())
		Expect(resp.Success).To(BeTrue())
	})

	When("name is empty", func() {
		It("should return validation error", func() {
			resp, err := srv.ServerGRPCClient().Create(ctx, &serverAPI.CreateRequest{
				Name: "",
				Type: serverAPI.ServerType_SERVER_TYPE_HTTP,
			})

			Expect(resp).To(BeNil())
			Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Name"))
		})
	})
})
