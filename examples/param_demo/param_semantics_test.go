package param_demo

import (
	"fmt"
	"testing"

	"github.com/eodnozerkin-ozon/allure-go-param/pkg/framework/provider"
	"github.com/eodnozerkin-ozon/allure-go-param/pkg/framework/suite"
)

// var testsCases = map[string][]string{
// 	"generatelistingid": {"all_value1"},
// 	"test1":             {"test1_valu1", "test1_valu1"},
// }

type GenerateListingId struct {
	suite.Suite
	//params map[string][]interface{}
	params int
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
	t.Title("Test1")
	t.WithNewStep("Test1 step1", func(ctx provider.StepCtx) {
		fmt.Println("Test1")
	})
}

func (s *GenerateListingId) Test2(t provider.T) {
	t.Title("Test2")
	t.WithNewStep("Test2 step1", func(ctx provider.StepCtx) {
		fmt.Println("Test2 step1")
	})
}

func TestGenerateListingId(t *testing.T) {
	t.Parallel()
	x := new(GenerateListingId)
	// m := make(map[string]interface{})
	//m["sdfsdsa"] = testsCases
	// x.params = map[string][]interface{}{
	// 	"all":   {"sdsad"},
	// 	"test1": {"dasxxz", "dffdsfd"},
	// }
	x.params = 5
	suite.RunSuite(t, x)
}
