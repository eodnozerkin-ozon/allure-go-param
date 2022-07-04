package manager

import (
	"github.com/eodnozerkin-ozon/allure-go-param/pkg/allure"
)

// Step adds step to test result
func (a *allureManager) Step(step *allure.Step) {
	a.ExecutionContext().AddStep(step)
}

// NewStep creates new step and adds it to test result
func (a *allureManager) NewStep(stepName string, params ...allure.Parameter) {
	a.ExecutionContext().AddStep(allure.NewSimpleStep(stepName, params...))
}
