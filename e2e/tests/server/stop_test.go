//go:build e2e

package server_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/status-mok/server/e2e/app"
	serverHTTPapi "github.com/status-mok/server/e2e/app/http-client/client/server_api/server_service"
	"github.com/status-mok/server/e2e/app/http-client/models"
	"github.com/status-mok/server/internal/pkg/tester"
	serverAPI "github.com/status-mok/server/pkg/server-api"
)

var _ = Describe("Stop method", Ordered, func() {
	var srv *app.TestAppServer
	serverNameForGRPC := "grpc-123"
	serverNameForHTTP := "http-123"

	BeforeAll(func() {
		srv = app.NewAppServer()

		_, err := srv.CreateRunningServers(ctx, serverNameForGRPC, serverNameForHTTP)
		Expect(err).To(BeNil())
	})

	AfterAll(func() {
		srv.Close()
	})

	Context("GRPC", func() {
		It("should finish successfully", func() {
			resp, err := srv.GRPCClient().ServerService().Stop(ctx, &serverAPI.StopRequest{
				Name: serverNameForGRPC,
			})

			Expect(err).To(BeNil())
			Expect(resp.Success).To(BeTrue())
		})

		When("app is already stopped", func() {
			It("should return a 'already stopped' error", func() {
				resp, err := srv.GRPCClient().ServerService().Stop(ctx, &serverAPI.StopRequest{
					Name: serverNameForGRPC,
				})

				Expect(resp).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("already stopped"))
			})
		})

		When("server does not exist", func() {
			It("should return a 'not found' error", func() {
				resp, err := srv.GRPCClient().ServerService().Stop(ctx, &serverAPI.StopRequest{
					Name: "not exist",
				})

				Expect(resp).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("not found"))
			})
		})

		Context("with request validation issues", func() {
			When("name is empty", func() {
				It("should return a validation error", func() {
					resp, err := srv.GRPCClient().ServerService().Stop(ctx, &serverAPI.StopRequest{
						Name: "",
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid StopRequest.Name"))
				})
			})
		})
	})

	Context("HTTP", func() {
		It("should finish successfully", func() {
			resp, err := srv.HTTPClient().ServerService().ServerServiceStop(&serverHTTPapi.ServerServiceStopParams{
				Body: &models.ServerServiceStopRequest{
					Name: tester.StringPtr(serverNameForHTTP),
				},
				Context: ctx,
			})

			Expect(err).To(BeNil())
			Expect(resp.IsSuccess()).To(BeTrue())
			Expect(resp.GetPayload().Success).To(BeTrue())
		})

		When("app is already stopped", func() {
			It("should return a 'already stopped' error", func() {
				resp, err := srv.HTTPClient().ServerService().ServerServiceStop(&serverHTTPapi.ServerServiceStopParams{
					Body: &models.ServerServiceStopRequest{
						Name: tester.StringPtr(serverNameForHTTP),
					},
					Context: ctx,
				})

				Expect(resp).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("already stopped"))
			})
		})

		When("server does not exist", func() {
			It("should return a 'not found' error", func() {
				resp, err := srv.HTTPClient().ServerService().ServerServiceStop(&serverHTTPapi.ServerServiceStopParams{
					Body: &models.ServerServiceStopRequest{
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
					resp, err := srv.HTTPClient().ServerService().ServerServiceStop(&serverHTTPapi.ServerServiceStopParams{
						Body: &models.ServerServiceStopRequest{
							Name: tester.StringPtr(""),
						},
						Context: ctx,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid StopRequest.Name"))
				})
			})
		})
	})
})
