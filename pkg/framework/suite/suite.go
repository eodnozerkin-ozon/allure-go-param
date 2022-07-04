package suite

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"regexp"

	"github.com/eodnozerkin-ozon/allure-go-param/pkg/framework/provider"
	"github.com/eodnozerkin-ozon/allure-go-param/pkg/framework/runner"
)

type InternalSuite interface {
	GetRunner() runner.TestRunner
	SetRunner(runner runner.TestRunner)
}

type Suite struct {
	runner runner.TestRunner
}

func (s *Suite) GetRunner() runner.TestRunner {
	return s.runner
}

func (s *Suite) SetRunner(runner runner.TestRunner) {
	s.runner = runner
}

func (s *Suite) RunSuite(t provider.T, suite InternalSuite) {
	t.SkipOnPrint()
	RunSuite(t.RealT(), suite)
}

func collectTests(runner *suiteRunner, suite InternalSuite) *suiteRunner {
	methodFinder := reflect.TypeOf(suite)
	for i := 0; i < methodFinder.NumMethod(); i++ {
		method := methodFinder.Method(i)

		ok, err := methodFilter(method.Name)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "allire-go: invalid regexp for -m: %s\n", err)
			os.Exit(1)
		}

		if !ok {
			continue
		}
		runner.AddTest(method.Name, method)
	}
	return runner
}

func collectHooks(runner *suiteRunner, suite InternalSuite) *suiteRunner {
	if beforeAll, ok := suite.(AllureBeforeSuite); ok {
		runner.BeforeAll(beforeAll.BeforeAll)
	}

	if beforeEach, ok := suite.(AllureBeforeTest); ok {
		runner.BeforeEach(beforeEach.BeforeEach)
	}

	if afterAll, ok := suite.(AllureAfterSuite); ok {
		runner.AfterAll(afterAll.AfterAll)
	}

	if afterEach, ok := suite.(AllureAfterTest); ok {
		runner.AfterEach(afterEach.AfterEach)
	}

	return runner
}

var matchMethod = flag.String("allure-go.m", "", "regular expression to select tests of the testify suite to run")

// Filtering method according to set regular expression
// specified command-line argument -m
func methodFilter(name string) (bool, error) {
	if ok, _ := regexp.MatchString("^Test", name); !ok {
		return false, nil
	}
	return regexp.MatchString(*matchMethod, name)
}

func RunSuite(t TestingT, suite InternalSuite) map[string]bool {
	return NewSuiteRunner(t, getPackage(2), t.Name(), suite).RunTests()
}

func RunNamedSuite(t TestingT, suiteName string, suite InternalSuite) map[string]bool {
	return NewSuiteRunner(t, getPackage(2), suiteName, suite).RunTests()
}
