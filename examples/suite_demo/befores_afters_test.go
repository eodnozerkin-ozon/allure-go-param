//go:build examples_new
// +build examples_new

package suite_demo

import (
	"testing"

	"github.com/eodnozerkin-ozon/allure-go-param/pkg/framework/provider"
	"github.com/eodnozerkin-ozon/allure-go-param/pkg/framework/suite"
)

type BeforeAfterDemoSuite struct {
	suite.Suite
}

func (s *BeforeAfterDemoSuite) BeforeEach(t provider.T) {
	t.NewStep("Before Test Step")
}

func (s *BeforeAfterDemoSuite) AfterEach(t provider.T) {
	t.NewStep("After Test Step")
}

func (s *BeforeAfterDemoSuite) BeforeAll(t provider.T) {
	t.NewStep("Before suite Step")
}

func (s *BeforeAfterDemoSuite) AfterAll(t provider.T) {
	t.NewStep("After suite Step")
}

func (s *BeforeAfterDemoSuite) TestBeforeAfterTest(t provider.T) {
	t.Epic("Demo")
	t.Feature("BeforeAfter")
	t.Title("Test wrapped with SetUp & TearDown")
	t.Description(`
		This test wrapped with SetUp and TearDown containert.`)

	t.Tags("BeforeAfter")
}

func TestBeforesAfters(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(BeforeAfterDemoSuite))
}
