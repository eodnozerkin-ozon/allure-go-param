package manager

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/eodnozerkin-ozon/allure-go-param/pkg/allure"
	"github.com/eodnozerkin-ozon/allure-go-param/pkg/framework/provider"
)

type testMetaMockLabels struct {
	result    *allure.Result
	container *allure.Container
	be        func(t provider.T)
	ae        func(t provider.T)
}

func (m *testMetaMockLabels) GetResult() *allure.Result {
	return m.result
}

func (m *testMetaMockLabels) SetResult(result *allure.Result) {
	m.result = result
}

func (m *testMetaMockLabels) GetContainer() *allure.Container {
	return m.container
}

func (m *testMetaMockLabels) SetBeforeEach(hook func(t provider.T)) {
	m.be = hook
}

func (m *testMetaMockLabels) GetBeforeEach() func(t provider.T) {
	return m.be
}

func (m *testMetaMockLabels) SetAfterEach(hook func(t provider.T)) {
	m.ae = hook
}

func (m *testMetaMockLabels) GetAfterEach() func(t provider.T) {
	return m.ae
}

func TestAllureManager_Labels(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockLabels{result: &allure.Result{}}}

	t.Run("ID", func(t *testing.T) {
		manager.ID("id")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.ID))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.ID), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.ID)[0].Value, "id")
	})

	t.Run("AllureID", func(t *testing.T) {
		manager.AllureID("allureID")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.AllureID))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.AllureID), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.AllureID)[0].Value, "allureID")
	})

	t.Run("Epic", func(t *testing.T) {
		manager.Epic("epic")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Epic))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Epic), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Epic)[0].Value, "epic")
	})

	t.Run("Feature", func(t *testing.T) {
		manager.Feature("feature")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Feature))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Feature), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Feature)[0].Value, "feature")
	})

	t.Run("Story", func(t *testing.T) {
		manager.Story("story")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Story))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Story), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Story)[0].Value, "story")
	})

	t.Run("Severity", func(t *testing.T) {
		manager.Severity(allure.TRIVIAL)
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Severity))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Severity), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Severity)[0].Value, allure.TRIVIAL.ToString())
	})

	t.Run("Tag", func(t *testing.T) {
		manager.Tag("tag1")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Tag))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Tag), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Tag)[0].Value, "tag1")
	})

	t.Run("Tags", func(t *testing.T) {
		manager.Tags("tag2", "tag3")
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Tag), 3)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Tag)[1].Value, "tag2")
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Tag)[2].Value, "tag3")
	})

	t.Run("Owner", func(t *testing.T) {
		manager.Owner("owner")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Owner))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Owner), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Owner)[0].Value, "owner")
	})

	t.Run("Lead", func(t *testing.T) {
		manager.Lead("lead")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Lead))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Lead), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Lead)[0].Value, "lead")
	})

	t.Run("Label", func(t *testing.T) {
		manager.Label(allure.NewLabel(allure.Framework, "Framework"))
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Framework))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Framework), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Framework)[0].Value, "Framework")
	})

	t.Run("Labels", func(t *testing.T) {
		manager.Labels(allure.NewLabel(allure.Tag, "tag4"), allure.NewLabel(allure.Tag, "tag5"))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Tag), 5)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Tag)[3].Value, "tag4")
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Tag)[4].Value, "tag5")
	})

	t.Run("AddSuiteLabel", func(t *testing.T) {
		manager.AddSuiteLabel("Suite")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Suite))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Suite), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Suite)[0].Value, "Suite")
	})

	t.Run("AddSubSuite", func(t *testing.T) {
		manager.AddSubSuite("SubSuite")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.SubSuite))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.SubSuite), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.SubSuite)[0].Value, "SubSuite")
	})

	t.Run("AddParentSuite", func(t *testing.T) {
		manager.AddParentSuite("ParentSuite")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.ParentSuite))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.ParentSuite), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.ParentSuite)[0].Value, "ParentSuite")
	})

	t.Run("ID", func(t *testing.T) {
		manager.Host("Host")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Host))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Host), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Host)[0].Value, "Host")
	})

	t.Run("Thread", func(t *testing.T) {
		manager.Thread("Thread")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Thread))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Thread), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Thread)[0].Value, "Thread")
	})

	t.Run("Language", func(t *testing.T) {
		manager.Language("Language")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Language))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Language), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Language)[0].Value, "Language")
	})

	t.Run("Package", func(t *testing.T) {
		manager.Package("Package")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Package))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Package), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Package)[0].Value, "Package")
	})

	t.Run("ReplaceLabel", func(t *testing.T) {
		manager.ReplaceLabel(allure.NewLabel(allure.Framework, "NewFramework"))
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Framework))
		require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Framework), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Framework)[0].Value, "NewFramework")
	})
}
