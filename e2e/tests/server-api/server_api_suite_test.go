//go:build e2e

package server_api_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestServerAPISuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server API Suite")
}
