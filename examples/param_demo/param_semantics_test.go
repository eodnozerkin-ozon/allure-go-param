package param_demo

import (
	"fmt"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	testInput string
	expected  string
}

var testsCases = map[string][]interface{}{
	"all": {testCase{testInput: "&test", expected: "test"}},
	"Test1": {
		testCase{testInput: "good", expected: "good"},
		testCase{testInput: " ev  il ", expected: "evil"},
		testCase{testInput: "", expected: ""},
	},
}

type GenerateListingId struct {
	suite.Suite

	params map[string][]interface{}
}

func (s *GenerateListingId) BeforeAll(t provider.T) {
	t.WithNewStep("BeforeAll step1", func(ctx provider.StepCtx) {
		fmt.Println("BeforeAll step1")
	})
}

func (s *GenerateListingId) BeforeEach(t provider.T) {
	t.WithNewStep("BeforeEach step1", func(ctx provider.StepCtx) {
		fmt.Println("BeforeEach step1")
	})
}

func (s *GenerateListingId) Test1(t provider.T) {
	param := t.GetParam().(testCase)
	t.Title("Test1")
	t.WithNewStep("Test1 step1", func(ctx provider.StepCtx) {
		actual := stripChars(param.testInput, " ")
		require.Equalf(
			t,
			param.expected,
			actual,
			"%s != %s",
			param.expected,
			actual,
		)
	})
}

func (s *GenerateListingId) Test2(t provider.T) {
	param := t.GetParam().(testCase)
	t.Title("Test2")
	t.WithNewStep("Test2 step1", func(ctx provider.StepCtx) {
		actual := stripChars(param.testInput, " ")
		require.Equalf(
			t,
			param.expected,
			actual,
			"%s != %s",
			param.expected,
			actual,
		)
	})
}

func TestGenerateListingId(t *testing.T) {
	x := new(GenerateListingId)
	x.params = testsCases

	suite.RunSuite(t, x)
}
