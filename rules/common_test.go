package rules

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

// Create a new test runner we'll use to test all the
// modules in the rules package
func TestRules(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rules Suite")
}

// Helper functions to create pointers to decimal values
func decimalPtr(decimal decimal.Decimal) *decimal.Decimal {
	return &decimal
}

// Helper functions to create pointers to time values
func timePtr(time time.Time) *time.Time {
	return &time
}
