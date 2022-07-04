package suite

import (
	"fmt"
	"os"
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/eodnozerkin-ozon/allure-go-param/pkg/framework/provider"
	"github.com/eodnozerkin-ozon/allure-go-param/pkg/framework/runner"
)

type suiteRunnerTMock struct {
	testing.TB

	t        *testing.T
	failNow  bool
	parallel bool
}

func (m *suiteRunnerTMock) Name() string {
	return "testSuite"
}

func (m *suiteRunnerTMock) Parallel() {
	m.parallel = true
}

func (m *suiteRunnerTMock) FailNow() {
	m.failNow = true
}

func (m *suiteRunnerTMock) Run(testName string, testBody func(t *testing.T)) bool {
	testBody(m.t)
	return true
}

func TestSuiteRunner_AddTest(t *testing.T) {
	r := suiteRunner{
		TestRunner: runner.NewRunner(t, "TestSuite"),
		tests:      map[string]*suiteTest{},
	}
	method := reflect.Method{
		Name:    "test",
		PkgPath: "",
		Type:    nil,
		Func:    reflect.Value{},
		Index:   0,
	}
	tags := []string{"tag1", "tag2"}
	testName := fmt.Sprintf("%s/%s", t.Name(), "test")
	r.AddTest("test", method, tags...)

	require.NotNil(t, r.tests[testName])
	require.Equal(t, "test", r.tests[testName].testName)
	require.Equal(t, method, r.tests[testName].testBody)
	require.Equal(t, tags, r.tests[testName].tags)
}

type TestSuiteRunner struct {
	Suite
	testSome1 bool
	testSome2 bool
}

func (s *TestSuiteRunner) TestSome1(t provider.T) {
	s.testSome1 = true
}

func (s *TestSuiteRunner) TestSome2(t provider.T) {
	s.testSome2 = true
}

func TestRunner_RunTests(t *testing.T) {
	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)

	suite := new(TestSuiteRunner)

	r := NewSuiteRunner(t, "packageName", "suiteName", suite)
	r.RunTests()

	require.True(t, suite.testSome1)
	require.True(t, suite.testSome2)
}

type TestSuiteRunnerPanic struct {
	Suite
	wg        sync.WaitGroup
	testSome1 bool
}

func (s *TestSuiteRunnerPanic) TestSome1(t provider.T) {
	defer s.wg.Done()
	s.testSome1 = true
	panic("whoops")
}

func TestRunner_RunTests_panic(t *testing.T) {
	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)

	suite := new(TestSuiteRunnerPanic)
	suite.wg = sync.WaitGroup{}
	mockT := &suiteRunnerTMock{t: new(testing.T)}
	r := NewSuiteRunner(mockT, "packageName", "suiteName", suite)
	suite.wg.Add(1)
	go require.NotPanics(t, func() {
		r.RunTests()
	})
	suite.wg.Wait()
	require.True(t, suite.testSome1)
}

type TestSuiteRunnerHooks struct {
	Suite
	wg        sync.WaitGroup
	testSome1 bool

	beforeAll  bool
	beforeEach bool
	afterEach  bool
	afterAll   bool
}

func (s *TestSuiteRunnerHooks) BeforeAll(t provider.T) {
	s.beforeAll = true
}

func (s *TestSuiteRunnerHooks) BeforeEach(t provider.T) {
	s.beforeEach = true
}

func (s *TestSuiteRunnerHooks) AfterEach(t provider.T) {
	s.afterEach = true
}

func (s *TestSuiteRunnerHooks) AfterAll(t provider.T) {
	s.afterAll = true
}

func (s *TestSuiteRunnerHooks) TestSome(t provider.T) {
	s.testSome1 = true
}

func TestRunner_hooks(t *testing.T) {
	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)

	suite := new(TestSuiteRunnerHooks)
	mockT := &suiteRunnerTMock{t: new(testing.T)}
	r := NewSuiteRunner(mockT, "packageName", "suiteName", suite)
	r.RunTests()

	require.True(t, suite.beforeAll)
	require.True(t, suite.beforeEach)
	require.True(t, suite.afterEach)
	require.True(t, suite.afterAll)
	require.True(t, suite.testSome1)
}
