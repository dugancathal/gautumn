package gautumn_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGautumn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gautumn Suite")
}
