package common

import (
	"sync"

	"github.com/eodnozerkin-ozon/allure-go-param/pkg/allure"
	"github.com/eodnozerkin-ozon/allure-go-param/pkg/framework/provider"
)

type ParentT interface {
	GetProvider() provider.Provider
	GetResult() *allure.Result
}

type HookProvider interface {
	BeforeEachContext()
	AfterEachContext()
	BeforeAllContext()
	AfterAllContext()

	GetSuiteMeta() provider.SuiteMeta
	GetTestMeta() provider.TestMeta
}

type InternalT interface {
	provider.T

	WG() *sync.WaitGroup
}
