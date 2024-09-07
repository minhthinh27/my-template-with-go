package biz

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"my-template-with-go/internal/entity"
	"my-template-with-go/response"
	"testing"
)

type MockArticleRepo struct {
	mock.Mock
}

func (m *MockArticleRepo) List() ([]*entity.Article, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Article), args.Error(1)
}

func (m *MockArticleRepo) Detail(id uint) (*entity.Article, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Article), args.Error(1)
}

func (m *MockArticleRepo) Create(item *entity.Article) error {
	return m.Called(item).Error(0)
}

func (m *MockArticleRepo) Update(id uint, items map[string]interface{}) error {
	return m.Called(id, items).Error(0)
}

func (m *MockArticleRepo) Delete(ids []uint) error {
	return m.Called(ids).Error(0)
}

func TestArticleUC_List(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockArticleRepo)

	// Create some fake data
	articles := []*entity.Article{
		{ID: 1, Title: "Test Title 1", Author: "Test Author 1"},
		{ID: 2, Title: "Test Title 2", Author: "Test Author 2"},
	}

	// Set up expected calls and return values
	mockRepo.On("List").Return(articles, nil)

	// Create the use case with the mock repository
	uc := NewArticleUseCase(mockRepo, nil)

	// Call the List method
	result, err := uc.List(echo.New().NewContext(nil, nil))

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result.([]*response.ArticleListRes)))
	assert.Equal(t, "Test Title 1", result.([]*response.ArticleListRes)[0].Title)

	// Ensure that the expectations were met
	mockRepo.AssertExpectations(t)
}
