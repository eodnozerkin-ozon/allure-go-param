package helper

import "github.com/ozontech/allure-go/pkg/framework/asserts_wrapper/wrapper"

// NewAssertsHelper inits new Assert interface
func NewAssertsHelper(t TestingT) AssertsHelper {
	return &a{
		t:       t,
		asserts: wrapper.NewAsserts(t),
	}
}
