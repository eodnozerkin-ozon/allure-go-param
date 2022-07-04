package helper

import (
	"github.com/eodnozerkin-ozon/allure-go-param/pkg/framework/asserts_wrapper/wrapper"
)

// NewRequireHelper inits new Require interface
func NewRequireHelper(t ProviderT) AssertsHelper {
	return &a{
		t:       t,
		asserts: wrapper.NewRequire(t),
	}
}
