package helper

import "github.com/eodnozerkin-ozon/allure-go-param/pkg/framework/asserts_wrapper/wrapper"

// NewAssertsHelper inits new Assert interface
func NewAssertsHelper(t ProviderT) AssertsHelper {
	return &a{
		t:       t,
		asserts: wrapper.NewAsserts(t),
	}
}
