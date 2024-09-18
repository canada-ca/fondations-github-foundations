package terragrunt

import (
	"errors"
	"fmt"
	"gh_foundations/internal/pkg/types/terraform_state/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"
)

type TeamTestSuite struct {
	suite.Suite
	mockStateExplorer *mocks.MockIStateExplorer
	resolver          TeamImportIdResolver
}

func (s *TeamTestSuite) SetupTest() {
	s.mockStateExplorer = new(mocks.MockIStateExplorer)
	s.resolver = TeamImportIdResolver{
		StateExplorer: s.mockStateExplorer,
	}
}

func TestTeamTestSuite(t *testing.T) {
	suite.Run(t, new(TeamTestSuite))
}

func (s *TeamTestSuite) TestTeamImportIdResolverResolveImportId() {
	resourceAddress := "some/resource/address"
	expectedString := "import-id"
	result := &gjson.Result{Type: gjson.String, Str: expectedString}

	s.mockStateExplorer.EXPECT().GetResourceChangeAfterAttribute(resourceAddress, "name").Return(result, nil)

	importID, err := s.resolver.ResolveImportId(resourceAddress)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedString, importID)
	s.mockStateExplorer.AssertExpectations(s.T())
}

func (s *TeamTestSuite) TestTeamImportIdResolverResolveImportIdGJsonErrorFailure() {
	resourceAddress := "some/resource/address"
	expectedString := ""
	expectedErr := errors.New("gjson error")

	s.mockStateExplorer.EXPECT().GetResourceChangeAfterAttribute(resourceAddress, "name").Return(nil, expectedErr)

	importID, err := s.resolver.ResolveImportId(resourceAddress)

	assert.Error(s.T(), err)
	assert.Equal(s.T(), expectedErr, err)
	assert.Equal(s.T(), expectedString, importID)
	s.mockStateExplorer.AssertExpectations(s.T())
}

func (s *TeamTestSuite) TestTeamImportIdResolverResolveImportIdKeyDoesNotExistFailure() {
	resourceAddress := "some/resource/address"
	expectedString := ""
	result := &gjson.Result{Type: gjson.Null}

	s.mockStateExplorer.EXPECT().GetResourceChangeAfterAttribute(resourceAddress, "name").Return(result, nil)

	importID, err := s.resolver.ResolveImportId(resourceAddress)

	assert.Error(s.T(), err)
	assert.Equal(s.T(), fmt.Sprintf("unable to resolve import id: missing %q attribute", "name"), err.Error())
	assert.Equal(s.T(), expectedString, importID)
	s.mockStateExplorer.AssertExpectations(s.T())
}

type TeamMemberTestSuite struct {
	suite.Suite
	mockStateExplorer *mocks.MockIStateExplorer
	resolver          TeamMemberImportIdResolver
}

func (s *TeamMemberTestSuite) SetupTest() {
	s.mockStateExplorer = new(mocks.MockIStateExplorer)
	s.resolver = TeamMemberImportIdResolver{
		StateExplorer: s.mockStateExplorer,
	}
}

func TestTeamMemberTestSuite(t *testing.T) {
	suite.Run(t, new(TeamMemberTestSuite))
}

func (s *TeamMemberTestSuite) TestTeamMemberImportIdResolverResolveImportId() {
	resourceAddress := "some/resource/address"
	teamIdResult := &gjson.Result{Type: gjson.String, Str: "team"}
	usernameResult := &gjson.Result{Type: gjson.String, Str: "username"}
	expectedImportID := "team:username"

	s.mockStateExplorer.EXPECT().GetResourceChangeAfterAttribute(resourceAddress, "team_id").Return(teamIdResult, nil).Once()
	s.mockStateExplorer.EXPECT().GetResourceChangeAfterAttribute(resourceAddress, "username").Return(usernameResult, nil).Once()

	importID, err := s.resolver.ResolveImportId(resourceAddress)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedImportID, importID)

	s.mockStateExplorer.AssertExpectations(s.T())
}
func (s *TeamMemberTestSuite) TestTeamMemberImportIdResolverResolveImportIdGjsonErrorFailure() {
	resourceAddress := "some/resource/address"
	teamIdResult := &gjson.Result{Type: gjson.String, Str: "team"}
	usernameResult := &gjson.Result{Type: gjson.String, Str: "username"}

	s.mockStateExplorer.EXPECT().GetResourceChangeAfterAttribute(resourceAddress, "team_id").Return(nil, errors.New("gjson error")).Once()
	s.mockStateExplorer.EXPECT().GetResourceChangeAfterAttribute(resourceAddress, "username").Return(usernameResult, nil).Once()

	s.mockStateExplorer.EXPECT().GetResourceChangeAfterAttribute(resourceAddress, "team_id").Return(teamIdResult, nil).Once()
	s.mockStateExplorer.EXPECT().GetResourceChangeAfterAttribute(resourceAddress, "username").Return(nil, errors.New("gjson error")).Once()

	s.mockStateExplorer.EXPECT().GetResourceChangeAfterAttribute(resourceAddress, "team_id").Return(nil, errors.New("gjson error")).Once()
	s.mockStateExplorer.EXPECT().GetResourceChangeAfterAttribute(resourceAddress, "username").Return(nil, errors.New("gjson error")).Once()

	expectedImportIDs := []string{":username", "team:", ":"}
	for _, expectedImportID := range expectedImportIDs {
		importID, err := s.resolver.ResolveImportId(resourceAddress)

		assert.Error(s.T(), err)
		assert.Equal(s.T(), expectedImportID, importID)
	}
	s.mockStateExplorer.AssertExpectations(s.T())
}
