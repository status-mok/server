//go:build e2e

package server_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/status-mok/server/e2e/server"
	serverAPI "github.com/status-mok/server/pkg/server-api"
)

var _ = Describe("Create method", Ordered, func() {
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

	Context("with request validation issues", func() {
		When("name is empty", func() {
			It("should return a validation error", func() {
				resp, err := srv.ServerGRPCClient().Create(ctx, &serverAPI.CreateRequest{
					Name: "",
					Type: serverAPI.ServerType_SERVER_TYPE_HTTP,
				})

				Expect(resp).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Name"))
			})
		})

		When("type is empty", func() {
			It("should return a validation error", func() {
				resp, err := srv.ServerGRPCClient().Create(ctx, &serverAPI.CreateRequest{
					Name: "123",
				})

				Expect(resp).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("unknown server type"))
			})
		})

		When("type is invalid", func() {
			It("should return a validation error", func() {
				resp, err := srv.ServerGRPCClient().Create(ctx, &serverAPI.CreateRequest{
					Name: "123",
					Type: -1,
				})

				Expect(resp).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Type"))
			})
		})
	})
})
