package terragrunt

import (
	"encoding/json"
	"errors"
	"gh_foundations/internal/pkg/types"
	typeMocks "gh_foundations/internal/pkg/types/mocks"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestTerragruntArchiveTestSuite(t *testing.T) {
	suite.Run(t, new(TerragruntArchiveTestSuite))
}

type TerragruntArchiveTestSuite struct {
	suite.Suite
	mockCmdExecutor *typeMocks.MockICommandExecutor
}

func (suite *TerragruntArchiveTestSuite) SetupTest() {
	fs = afero.NewMemMapFs()
	suite.mockCmdExecutor = new(typeMocks.MockICommandExecutor)
}

func (suite *TerragruntArchiveTestSuite) TestNewTerragruntPlanFile() {
	name := "test"
	modulePath := "path/to/module"
	moduleDir := "path/to/module/dir"
	outputFilePath := "path/to/output/file"

	planFile, err := NewTerragruntPlanFile(name, modulePath, moduleDir, outputFilePath)

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), planFile)
	assert.Equal(suite.T(), name, planFile.Name)
	assert.Equal(suite.T(), modulePath, planFile.ModulePath)
	assert.Equal(suite.T(), moduleDir, planFile.ModuleDir)
	assert.Equal(suite.T(), outputFilePath, planFile.OutputFilePath)
}

func (suite *TerragruntArchiveTestSuite) TestNewTerragruntPlanFileAlreadyExists() {
	name := "test"
	modulePath := "path/to/module"
	moduleDir := "path/to/module/dir"
	outputFilePaths := []string{"path/to/output/file", "path/to/output/file.json"}
	expectedOutputFilePaths := []string{"path/to/output/copy_file", "path/to/output/copy_file.json"}

	for i := range outputFilePaths {
		f, err := fs.Create(outputFilePaths[i])
		require.NoError(suite.T(), err)
		f.Close()

		planFile, err := NewTerragruntPlanFile(name, modulePath, moduleDir, outputFilePaths[i])

		assert.NoError(suite.T(), err)
		assert.NotNil(suite.T(), planFile)
		assert.Equal(suite.T(), name, planFile.Name)
		assert.Equal(suite.T(), modulePath, planFile.ModulePath)
		assert.Equal(suite.T(), moduleDir, planFile.ModuleDir)
		assert.Equal(suite.T(), expectedOutputFilePaths[i], planFile.OutputFilePath)
	}
}

func (suite *TerragruntArchiveTestSuite) TestPlanFileRunPlan() {
	actualArgs := make([]string, 0)
	newCommandExecutor = func(_ string, args ...string) types.ICommandExecutor {
		if args[0] == "plan" {
			actualArgs = args
		}
		return suite.mockCmdExecutor
	}
	planFile := &PlanFile{
		Name:           "test",
		ModulePath:     "path/to/module",
		ModuleDir:      "path/to/module/dir",
		OutputFilePath: "path/to/output/file",
	}

	suite.mockCmdExecutor.EXPECT().Run().Return(nil)
	suite.mockCmdExecutor.EXPECT().SetDir("path/to/module/dir").Return()
	suite.mockCmdExecutor.EXPECT().SetOutput(mock.Anything).Return()
	suite.mockCmdExecutor.EXPECT().SetErrorOutput(mock.Anything).Return()
	suite.mockCmdExecutor.EXPECT().String().Return("")

	targets := []string{"", "test-target"}
	for _, target := range targets {
		var err error
		if target == "" {
			err = planFile.RunPlan(nil)
		} else {
			err = planFile.RunPlan(&target)
		}
		assert.NoError(suite.T(), err)
		if target != "" {
			assert.Contains(suite.T(), actualArgs, "-target="+target)
		}
	}
}

func (suite *TerragruntArchiveTestSuite) TestPlanFileRunPlanCommandFailure() {
	var errBufferWriterFunc func(errMessage string)
	newCommandExecutor = func(_ string, args ...string) types.ICommandExecutor {
		if args[0] == "plan" {
			suite.mockCmdExecutor.EXPECT().Run().RunAndReturn(func() error {
				errBufferWriterFunc("error bad plan")
				return errors.New("")
			}).Once()
		} else {
			suite.mockCmdExecutor.EXPECT().Run().Return(nil).Once()
		}
		return suite.mockCmdExecutor
	}
	planFile := &PlanFile{
		Name:           "test",
		ModulePath:     "path/to/module",
		ModuleDir:      "path/to/module/dir",
		OutputFilePath: "path/to/output/file",
	}
	expectedErrorMessage := "error running plan: error bad plan"

	suite.mockCmdExecutor.EXPECT().SetDir("path/to/module/dir").Return()
	suite.mockCmdExecutor.EXPECT().SetOutput(mock.Anything).Return()
	suite.mockCmdExecutor.EXPECT().SetErrorOutput(mock.Anything).Run(func(writer io.Writer) {
		errBufferWriterFunc = func(errMessage string) {
			writer.Write([]byte(errMessage))
		}
	}).Return()
	suite.mockCmdExecutor.EXPECT().String().Return("")

	err := planFile.RunPlan(nil)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedErrorMessage, err.Error())
}

func (suite *TerragruntArchiveTestSuite) TestPlanFileRunPlanFileCreateFailure() {
	fs = afero.NewReadOnlyFs(fs)
	newCommandExecutor = func(_ string, _ ...string) types.ICommandExecutor {
		return suite.mockCmdExecutor
	}
	planFile := &PlanFile{
		Name:           "test",
		ModulePath:     "path/to/module",
		ModuleDir:      "path/to/module/dir",
		OutputFilePath: "bad/path/plan.json",
	}

	// Expect to see error thrown by the afero read only file system
	expectedErrorMessage := "operation not permitted"

	suite.mockCmdExecutor.EXPECT().Run().Return(nil)
	suite.mockCmdExecutor.EXPECT().SetDir("path/to/module/dir").Return()
	suite.mockCmdExecutor.EXPECT().SetOutput(mock.Anything).Return()
	suite.mockCmdExecutor.EXPECT().SetErrorOutput(mock.Anything).Return()
	suite.mockCmdExecutor.EXPECT().String().Return("")

	err := planFile.RunPlan(nil)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedErrorMessage, err.Error())
}

func (suite *TerragruntArchiveTestSuite) TestPlanFileRunShowCommandFailure() {
	var errBufferWriterFunc func(errMessage string)
	newCommandExecutor = func(_ string, args ...string) types.ICommandExecutor {
		if args[0] == "show" {
			suite.mockCmdExecutor.EXPECT().Run().RunAndReturn(func() error {
				errBufferWriterFunc("error show failed")
				return errors.New("")
			}).Once()
		} else {
			suite.mockCmdExecutor.EXPECT().Run().Return(nil).Once()
		}
		return suite.mockCmdExecutor
	}
	planFile := &PlanFile{
		Name:           "test",
		ModulePath:     "path/to/module",
		ModuleDir:      "path/to/module/dir",
		OutputFilePath: "path/to/output/file",
	}
	expectedErrorMessage := "error outputting plan: error show failed"

	suite.mockCmdExecutor.EXPECT().SetDir("path/to/module/dir").Return()
	suite.mockCmdExecutor.EXPECT().SetOutput(mock.Anything).Return()
	suite.mockCmdExecutor.EXPECT().SetErrorOutput(mock.AnythingOfType("*bytes.Buffer")).Run(func(writer io.Writer) {
		errBufferWriterFunc = func(errMessage string) {
			writer.Write([]byte(errMessage))
		}
	}).Return()
	suite.mockCmdExecutor.EXPECT().String().Return("")

	err := planFile.RunPlan(nil)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedErrorMessage, err.Error())
}

func (suite *TerragruntArchiveTestSuite) TestPlanFileGetStateExplorer() {
	planFile := &PlanFile{
		Name:           "test",
		ModulePath:     "path/to/module",
		ModuleDir:      "path/to/module/dir",
		OutputFilePath: "plan.json",
	}
	versions := []string{"1.2"}
	for _, version := range versions {
		var fileName = "plan.json"
		jsonContents := map[string]any{
			"format_version": version,
		}
		bytes, err := json.Marshal(jsonContents)
		require.NoError(suite.T(), err)
		err = afero.WriteFile(fs, fileName, bytes, 0644)
		require.NoError(suite.T(), err)

		stateExplorer, err := planFile.GetStateExplorer()
		assert.NotNil(suite.T(), stateExplorer)
		assert.NoError(suite.T(), err)

		fs.Remove(fileName)
	}
}

func (suite *TerragruntArchiveTestSuite) TestPlanFileGetStateExplorerReadFileFailure() {
	planFile := &PlanFile{
		Name:           "test",
		ModulePath:     "path/to/module",
		ModuleDir:      "path/to/module/dir",
		OutputFilePath: "plan.json",
	}
	stateExplorer, err := planFile.GetStateExplorer()
	assert.Nil(suite.T(), stateExplorer)
	assert.Error(suite.T(), err)
}

func (suite *TerragruntArchiveTestSuite) TestPlanFileGetStateExplorerVersionQueryFailure() {
	planFile := &PlanFile{
		Name:           "test",
		ModulePath:     "path/to/module",
		ModuleDir:      "path/to/module/dir",
		OutputFilePath: "plan.json",
	}

	fileContents := []map[string]any{
		{
			"bad_format_version_key": 1.2,
		},
		{
			"format_version": map[string]any{
				"bad_format_version_type": 1.2,
			},
		},
	}

	for _, fileContent := range fileContents {
		var fileName = "plan.json"
		bytes, err := json.Marshal(fileContent)
		require.NoError(suite.T(), err)
		err = afero.WriteFile(fs, fileName, bytes, 0644)
		require.NoError(suite.T(), err)

		stateExplorer, err := planFile.GetStateExplorer()

		assert.Nil(suite.T(), stateExplorer)
		assert.Error(suite.T(), err)

		fs.Remove(fileName)
	}
}

func (suite *TerragruntArchiveTestSuite) TestPlanFileGetStateExplorerUnsupportedVersionFailure() {
	planFile := &PlanFile{
		Name:           "test",
		ModulePath:     "path/to/module",
		ModuleDir:      "path/to/module/dir",
		OutputFilePath: "plan.json",
	}

	var fileName = "plan.json"
	jsonContents := map[string]any{
		"format_version": "1.0",
	}
	bytes, err := json.Marshal(jsonContents)
	require.NoError(suite.T(), err)
	err = afero.WriteFile(fs, fileName, bytes, 0644)
	require.NoError(suite.T(), err)

	stateExplorer, err := planFile.GetStateExplorer()
	assert.Nil(suite.T(), stateExplorer)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "unsupported version \"1.0\"", err.Error())

	fs.Remove(fileName)
}
