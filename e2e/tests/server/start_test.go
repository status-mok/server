//go:build e2e

package server_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/status-mok/server/e2e/app"
	serverHTTPapi "github.com/status-mok/server/e2e/app/http-client/client/server_api/server_service"
	"github.com/status-mok/server/e2e/app/http-client/models"
	"github.com/status-mok/server/internal/pkg/tester"
	serverAPI "github.com/status-mok/server/pkg/server-api"
)

var _ = Describe("Start method", Ordered, func() {
	var srv *app.TestAppServer
	serverNameForGRPC := "grpc-123"
	serverNameForHTTP := "http-123"

	BeforeAll(func() {
		srv = app.NewAppServer()

		err := srv.CreateStoppedServers(ctx, serverNameForGRPC, serverNameForHTTP)
		Expect(err).To(BeNil())
	})

	AfterAll(func() {
		srv.Close()
	})

	Context("GRPC", func() {
		It("should finish successfully", func() {
			resp, err := srv.GRPCClient().ServerService().Start(ctx, &serverAPI.StartRequest{
				Name: serverNameForGRPC,
			})

			Expect(err).To(BeNil())
			Expect(resp.Success).To(BeTrue())
			address := resp.GetAddress()
			port := address[strings.LastIndex(address, ":")+1:]
			Expect(port).NotTo(Equal(0))
		})

		When("app is already running", func() {
			It("should return a 'already running' error", func() {
				resp, err := srv.GRPCClient().ServerService().Start(ctx, &serverAPI.StartRequest{
					Name: serverNameForGRPC,
				})

				Expect(resp).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("already running"))
			})
		})

		When("app does not exist", func() {
			It("should return a 'not found' error", func() {
				resp, err := srv.GRPCClient().ServerService().Start(ctx, &serverAPI.StartRequest{
					Name: "not exist",
				})

				Expect(resp).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("not found"))
			})
		})

		Context("with request validation issues", func() {
			When("name is empty", func() {
				It("should return a validation error", func() {
					resp, err := srv.GRPCClient().ServerService().Start(ctx, &serverAPI.StartRequest{
						Name: "",
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid StartRequest.Name"))
				})
			})
		})
	})

	Context("HTTP", func() {
		It("should finish successfully", func() {
			resp, err := srv.HTTPClient().ServerService().ServerServiceStart(&serverHTTPapi.ServerServiceStartParams{
				Body: &models.ServerServiceStartRequest{
					Name: tester.StringPtr(serverNameForHTTP),
				},
				Context: ctx,
			})

			Expect(err).To(BeNil())
			Expect(resp.IsSuccess()).To(BeTrue())
			Expect(resp.GetPayload().Success).To(BeTrue())
			address := resp.GetPayload().Address
			port := address[strings.LastIndex(address, ":")+1:]
			Expect(port).NotTo(Equal(0))
		})

		When("app is already running", func() {
			It("should return a 'already running' error", func() {
				resp, err := srv.HTTPClient().ServerService().ServerServiceStart(&serverHTTPapi.ServerServiceStartParams{
					Body: &models.ServerServiceStartRequest{
						Name: tester.StringPtr(serverNameForGRPC),
					},
					Context: ctx,
				})

				Expect(resp).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("already running"))
			})
		})

		When("app does not exist", func() {
			It("should return a 'not found' error", func() {
				resp, err := srv.HTTPClient().ServerService().ServerServiceStart(&serverHTTPapi.ServerServiceStartParams{
					Body: &models.ServerServiceStartRequest{
						Name: tester.StringPtr("not exist"),
					},
					Context: ctx,
				})

				Expect(resp).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("not found"))
			})
		})

		Context("with request validation issues", func() {
			When("name is empty", func() {
				It("should return a validation error", func() {
					resp, err := srv.HTTPClient().ServerService().ServerServiceStart(&serverHTTPapi.ServerServiceStartParams{
						Body: &models.ServerServiceStartRequest{
							Name: tester.StringPtr(""),
						},
						Context: ctx,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid StartRequest.Name"))
				})
			})
		})
	})
})
