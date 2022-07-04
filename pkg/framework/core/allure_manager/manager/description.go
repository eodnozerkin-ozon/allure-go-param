package manager

import (
	"github.com/eodnozerkin-ozon/allure-go-param/pkg/allure"
)

// Title changes default test name to title(string)
func (a *allureManager) Title(title string) {
	a.safely(func(result *allure.Result) {
		result.Name = title
	})
}

// Description provides description to test result
func (a *allureManager) Description(description string) {
	a.safely(func(result *allure.Result) {
		result.Description = description
	})
}
