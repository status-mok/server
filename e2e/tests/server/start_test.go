//go:build e2e

package server_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/status-mok/server/e2e/server"
	serverAPI "github.com/status-mok/server/pkg/server-api"
)

var _ = Describe("Start method", Ordered, func() {
	var srv *server.TestServer
	serverName := "123"

	BeforeAll(func() {
		srv = server.NewServer()
		resp, err := srv.ServerGRPCClient().Create(ctx, &serverAPI.CreateRequest{
			Name: serverName,
			Type: serverAPI.ServerType_SERVER_TYPE_HTTP,
		})
		Expect(err).To(BeNil())
		Expect(resp.Success).To(BeTrue())
	})

	AfterAll(func() {
		srv.Close()
	})

	It("should finish successfully", func() {
		resp, err := srv.ServerGRPCClient().Start(ctx, &serverAPI.StartRequest{
			Name: serverName,
		})

		Expect(err).To(BeNil())
		Expect(resp.Success).To(BeTrue())
	})

	When("server is already running", func() {
		It("should return a 'not found' error", func() {
			resp, err := srv.ServerGRPCClient().Start(ctx, &serverAPI.StartRequest{
				Name: serverName,
			})

			Expect(resp).To(BeNil())
			Expect(err.Error()).To(ContainSubstring("already running"))
		})
	})

	When("server does not exist", func() {
		It("should return a 'not found' error", func() {
			resp, err := srv.ServerGRPCClient().Start(ctx, &serverAPI.StartRequest{
				Name: "not exist",
			})

			Expect(resp).To(BeNil())
			Expect(err.Error()).To(ContainSubstring("not found"))
		})
	})

	Context("with request validation issues", func() {
		When("name is empty", func() {
			It("should return a validation error", func() {
				resp, err := srv.ServerGRPCClient().Start(ctx, &serverAPI.StartRequest{
					Name: "",
				})

				Expect(resp).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("invalid StartRequest.Name"))
			})
		})
	})
})
