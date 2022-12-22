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

var _ = Describe("Delete method", Ordered, func() {
	var srv *app.TestAppServer
	serverNameForGRPC := "grpc-123"
	serverNameForHTTP := "http-123"

	BeforeAll(func() {
		srv = app.NewAppServer()

		resp, err := srv.GRPCClient().ServerService().Create(ctx, &serverAPI.CreateRequest{
			Name: serverNameForGRPC,
			Type: serverAPI.ServerType_SERVER_TYPE_HTTP,
		})
		Expect(err).To(BeNil())
		Expect(resp.Success).To(BeTrue())

		resp, err = srv.GRPCClient().ServerService().Create(ctx, &serverAPI.CreateRequest{
			Name: serverNameForHTTP,
			Type: serverAPI.ServerType_SERVER_TYPE_HTTP,
		})
		Expect(err).To(BeNil())
		Expect(resp.Success).To(BeTrue())
	})

	AfterAll(func() {
		srv.Close()
	})

	Context("GRPC", func() {
		It("should finish successfully", func() {
			resp, err := srv.GRPCClient().ServerService().Delete(ctx, &serverAPI.DeleteRequest{
				Name: serverNameForGRPC,
			})

			Expect(err).To(BeNil())
			Expect(resp.Success).To(BeTrue())
		})

		When("app does not exist", func() {
			It("should return a 'not found' error", func() {
				resp, err := srv.GRPCClient().ServerService().Delete(ctx, &serverAPI.DeleteRequest{
					Name: serverNameForGRPC,
				})

				Expect(resp).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("not found"))
			})
		})

		Context("with request validation issues", func() {
			When("name is empty", func() {
				It("should return a validation error", func() {
					resp, err := srv.GRPCClient().ServerService().Delete(ctx, &serverAPI.DeleteRequest{
						Name: "",
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid DeleteRequest.Name"))
				})
			})
		})
	})

	Context("HTTP", func() {
		It("should finish successfully", func() {
			resp, err := srv.HTTPClient().ServerService().ServerServiceDelete(&serverHTTPapi.ServerServiceDeleteParams{
				Body: &models.ServerServiceDeleteRequest{
					Name: tester.StringPtr(serverNameForHTTP),
				},
				Context: ctx,
			})

			Expect(err).To(BeNil())
			Expect(resp.IsSuccess()).To(BeTrue())
			Expect(resp.GetPayload().Success).To(BeTrue())
		})

		When("app does not exist", func() {
			It("should return a 'not found' error", func() {
				resp, err := srv.HTTPClient().ServerService().ServerServiceDelete(&serverHTTPapi.ServerServiceDeleteParams{
					Body: &models.ServerServiceDeleteRequest{
						Name: tester.StringPtr(serverNameForHTTP),
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
					resp, err := srv.HTTPClient().ServerService().ServerServiceDelete(&serverHTTPapi.ServerServiceDeleteParams{
						Body: &models.ServerServiceDeleteRequest{
							Name: tester.StringPtr(""),
						},
						Context: ctx,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid DeleteRequest.Name"))
				})
			})
		})
	})
})
