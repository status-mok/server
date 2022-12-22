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

var _ = Describe("Create method", Ordered, func() {
	var srv *app.TestAppServer
	serverNameForGRPC := "grpc-123"
	serverNameForHTTP := "http-123"

	BeforeAll(func() {
		srv = app.NewAppServer()
	})

	AfterAll(func() {
		srv.Close()
	})

	Context("GRPC", func() {
		It("should finish successfully", func() {
			resp, err := srv.GRPCClient().ServerService().Create(ctx, &serverAPI.CreateRequest{
				Name: serverNameForGRPC,
				Type: serverAPI.ServerType_SERVER_TYPE_HTTP,
			})

			Expect(err).To(BeNil())
			Expect(resp.Success).To(BeTrue())
		})

		Context("with request validation issues", func() {
			When("name is empty", func() {
				It("should return a validation error", func() {
					resp, err := srv.GRPCClient().ServerService().Create(ctx, &serverAPI.CreateRequest{
						Name: "",
						Type: serverAPI.ServerType_SERVER_TYPE_HTTP,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Name"))
				})
			})

			When("port is invalid", func() {
				It("should return a validation error", func() {
					resp, err := srv.GRPCClient().ServerService().Create(ctx, &serverAPI.CreateRequest{
						Name: serverNameForGRPC,
						Port: 100_000,
						Type: serverAPI.ServerType_SERVER_TYPE_HTTP,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Port"))
				})
			})

			When("type is empty", func() {
				It("should return a validation error", func() {
					resp, err := srv.GRPCClient().ServerService().Create(ctx, &serverAPI.CreateRequest{
						Name: serverNameForGRPC,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("unknown server type"))
				})
			})

			When("type is invalid", func() {
				It("should return a validation error", func() {
					resp, err := srv.GRPCClient().ServerService().Create(ctx, &serverAPI.CreateRequest{
						Name: serverNameForGRPC,
						Type: -1,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Type"))
				})
			})
		})
	})

	Context("HTTP", func() {
		It("should finish successfully", func() {
			resp, err := srv.HTTPClient().ServerService().ServerServiceCreate(&serverHTTPapi.ServerServiceCreateParams{
				Body: &models.ServerServiceCreateRequest{
					Name: tester.StringPtr(serverNameForHTTP),
					Port: tester.Int64Ptr(0),
					Type: models.ServerServiceServerTypeSERVERTYPEHTTP.Pointer(),
				},
				Context: ctx,
			})

			Expect(err).To(BeNil())
			Expect(resp.IsSuccess()).To(BeTrue())
			Expect(resp.GetPayload().Success).To(BeTrue())
		})

		Context("with request validation issues", func() {
			When("name is empty", func() {
				It("should return a validation error", func() {
					resp, err := srv.HTTPClient().ServerService().ServerServiceCreate(&serverHTTPapi.ServerServiceCreateParams{
						Body: &models.ServerServiceCreateRequest{
							Name: tester.StringPtr(""),
							Port: tester.Int64Ptr(0),
							Type: models.ServerServiceServerTypeSERVERTYPEHTTP.Pointer(),
						},
						Context: ctx,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Name"))
				})
			})

			When("port is invalid", func() {
				It("should return a validation error", func() {
					resp, err := srv.HTTPClient().ServerService().ServerServiceCreate(&serverHTTPapi.ServerServiceCreateParams{
						Body: &models.ServerServiceCreateRequest{
							Name: tester.StringPtr(serverNameForHTTP),
							Port: tester.Int64Ptr(100_000),
							Type: models.ServerServiceServerTypeSERVERTYPEHTTP.Pointer(),
						},
						Context: ctx,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Port"))
				})
			})

			When("type is empty", func() {
				It("should return a validation error", func() {
					resp, err := srv.HTTPClient().ServerService().ServerServiceCreate(&serverHTTPapi.ServerServiceCreateParams{
						Body: &models.ServerServiceCreateRequest{
							Name: tester.StringPtr(serverNameForHTTP),
							Port: tester.Int64Ptr(0),
						},
						Context: ctx,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("unknown server type"))
				})
			})

			When("type is invalid", func() {
				It("should return a validation error", func() {
					resp, err := srv.HTTPClient().ServerService().ServerServiceCreate(&serverHTTPapi.ServerServiceCreateParams{
						Body: &models.ServerServiceCreateRequest{
							Name: tester.StringPtr(serverNameForHTTP),
							Port: tester.Int64Ptr(0),
							Type: models.ServerServiceServerType("unknown type").Pointer(),
						},
						Context: ctx,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid value for enum type"))
				})
			})
		})
	})
})
