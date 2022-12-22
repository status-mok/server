//go:build e2e

package route_test

import (
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/status-mok/server/e2e/app"
	routeHTTPapi "github.com/status-mok/server/e2e/app/http-client/client/route_api/route_service"
	"github.com/status-mok/server/e2e/app/http-client/models"
	"github.com/status-mok/server/internal/pkg/tester"
	routeAPI "github.com/status-mok/server/pkg/route-api"
)

var _ = Describe("Create method", Ordered, func() {
	var (
		srv     *app.TestAppServer
		addrMap map[string]string
		err     error
	)
	serverNameForGRPC := "grpc-123"
	serverNameForHTTP := "http-123"
	sampleURL := "/some-url"

	BeforeAll(func() {
		srv = app.NewAppServer()

		addrMap, err = srv.CreateRunningServers(ctx, serverNameForGRPC, serverNameForHTTP)
		Expect(err).To(BeNil())
	})

	AfterAll(func() {
		srv.Close()
	})

	Context("GRPC", func() {
		It("should finish successfully", func() {
			routeResp, err := http.Get(httpScheme + addrMap[serverNameForGRPC] + sampleURL)
			Expect(err).To(BeNil())
			Expect(routeResp.StatusCode).To(Equal(http.StatusNotFound))

			resp, err := srv.GRPCClient().RouteService().Create(ctx, &routeAPI.CreateRequest{
				ServerName: serverNameForGRPC,
				Url:        sampleURL,
				Type:       routeAPI.RouteType_ROUTE_TYPE_REQ_RESP,
			})

			Expect(err).To(BeNil())
			Expect(resp.Success).To(BeTrue())

			routeResp, err = http.Get(httpScheme + addrMap[serverNameForGRPC] + sampleURL)
			Expect(err).To(BeNil())
			Expect(routeResp.StatusCode).To(Equal(http.StatusOK))
		})

		Context("with request validation issues", func() {
			When("server name is empty", func() {
				It("should return a validation error", func() {
					resp, err := srv.GRPCClient().RouteService().Create(ctx, &routeAPI.CreateRequest{
						ServerName: "",
						Url:        sampleURL,
						Type:       routeAPI.RouteType_ROUTE_TYPE_REQ_RESP,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.ServerName"))
				})
			})

			When("url is empty", func() {
				It("should return a validation error", func() {
					resp, err := srv.GRPCClient().RouteService().Create(ctx, &routeAPI.CreateRequest{
						ServerName: serverNameForGRPC,
						Url:        "",
						Type:       routeAPI.RouteType_ROUTE_TYPE_REQ_RESP,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Url"))
				})
			})

			When("not relative url is invalid", func() {
				It("should return a validation error", func() {
					resp, err := srv.GRPCClient().RouteService().Create(ctx, &routeAPI.CreateRequest{
						ServerName: serverNameForGRPC,
						Url:        "asd",
						Type:       routeAPI.RouteType_ROUTE_TYPE_REQ_RESP,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Url"))
				})
			})

			When("url with GET params is invalid", func() {
				It("should return a validation error", func() {
					resp, err := srv.GRPCClient().RouteService().Create(ctx, &routeAPI.CreateRequest{
						ServerName: serverNameForGRPC,
						Url:        sampleURL + "?asdasd",
						Type:       routeAPI.RouteType_ROUTE_TYPE_REQ_RESP,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Url"))
				})
			})

			When("type is empty", func() {
				It("should return a validation error", func() {
					resp, err := srv.GRPCClient().RouteService().Create(ctx, &routeAPI.CreateRequest{
						ServerName: serverNameForGRPC,
						Url:        sampleURL,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("unknown route type"))
				})
			})

			When("type is invalid", func() {
				It("should return a validation error", func() {
					resp, err := srv.GRPCClient().RouteService().Create(ctx, &routeAPI.CreateRequest{
						ServerName: serverNameForGRPC,
						Url:        sampleURL,
						Type:       -1,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Type"))
				})
			})
		})
	})

	Context("HTTP", func() {
		It("should finish successfully", func() {
			routeResp, err := http.Get(httpScheme + addrMap[serverNameForHTTP] + sampleURL)
			Expect(err).To(BeNil())
			Expect(routeResp.StatusCode).To(Equal(http.StatusNotFound))

			resp, err := srv.HTTPClient().RouteService().RouteServiceCreate(&routeHTTPapi.RouteServiceCreateParams{
				Body: &models.RouteServiceCreateRequest{
					ServerName: tester.StringPtr(serverNameForHTTP),
					Type:       models.RouteServiceRouteTypeROUTETYPEREQRESP.Pointer(),
					URL:        tester.StringPtr(sampleURL),
				},
				Context: ctx,
			})

			Expect(err).To(BeNil())
			Expect(resp.IsSuccess()).To(BeTrue())
			Expect(resp.GetPayload().Success).To(BeTrue())

			routeResp, err = http.Get(httpScheme + addrMap[serverNameForHTTP] + sampleURL)
			Expect(err).To(BeNil())
			Expect(routeResp.StatusCode).To(Equal(http.StatusOK))
		})

		Context("with request validation issues", func() {
			When("server name is empty", func() {
				It("should return a validation error", func() {
					resp, err := srv.HTTPClient().RouteService().RouteServiceCreate(&routeHTTPapi.RouteServiceCreateParams{
						Body: &models.RouteServiceCreateRequest{
							ServerName: tester.StringPtr(""),
							Type:       models.RouteServiceRouteTypeROUTETYPEREQRESP.Pointer(),
							URL:        tester.StringPtr(sampleURL),
						},
						Context: ctx,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.ServerName"))
				})
			})

			When("url is empty", func() {
				It("should return a validation error", func() {
					resp, err := srv.HTTPClient().RouteService().RouteServiceCreate(&routeHTTPapi.RouteServiceCreateParams{
						Body: &models.RouteServiceCreateRequest{
							ServerName: tester.StringPtr(serverNameForHTTP),
							Type:       models.RouteServiceRouteTypeROUTETYPEREQRESP.Pointer(),
							URL:        tester.StringPtr(""),
						},
						Context: ctx,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Url"))
				})
			})

			When("not relative url is invalid", func() {
				It("should return a validation error", func() {
					resp, err := srv.HTTPClient().RouteService().RouteServiceCreate(&routeHTTPapi.RouteServiceCreateParams{
						Body: &models.RouteServiceCreateRequest{
							ServerName: tester.StringPtr(serverNameForHTTP),
							Type:       models.RouteServiceRouteTypeROUTETYPEREQRESP.Pointer(),
							URL:        tester.StringPtr("asd"),
						},
						Context: ctx,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Url"))
				})
			})

			When("url with GET params is invalid", func() {
				It("should return a validation error", func() {
					resp, err := srv.HTTPClient().RouteService().RouteServiceCreate(&routeHTTPapi.RouteServiceCreateParams{
						Body: &models.RouteServiceCreateRequest{
							ServerName: tester.StringPtr(serverNameForHTTP),
							Type:       models.RouteServiceRouteTypeROUTETYPEREQRESP.Pointer(),
							URL:        tester.StringPtr(sampleURL + "?asdasd"),
						},
						Context: ctx,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("invalid CreateRequest.Url"))
				})
			})

			When("type is empty", func() {
				It("should return a validation error", func() {
					resp, err := srv.HTTPClient().RouteService().RouteServiceCreate(&routeHTTPapi.RouteServiceCreateParams{
						Body: &models.RouteServiceCreateRequest{
							ServerName: tester.StringPtr(serverNameForHTTP),
							URL:        tester.StringPtr(sampleURL),
						},
						Context: ctx,
					})

					Expect(resp).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("unknown route type"))
				})
			})

			When("type is invalid", func() {
				It("should return a validation error", func() {
					resp, err := srv.HTTPClient().RouteService().RouteServiceCreate(&routeHTTPapi.RouteServiceCreateParams{
						Body: &models.RouteServiceCreateRequest{
							ServerName: tester.StringPtr(serverNameForHTTP),
							Type:       models.RouteServiceRouteType("unknown type").Pointer(),
							URL:        tester.StringPtr(sampleURL),
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
