// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "ref/entity"
	reflect "reflect"
)

// MockIngredientRepository is a mock of IngredientRepository interface
type MockIngredientRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIngredientRepositoryMockRecorder
}

// MockIngredientRepositoryMockRecorder is the mock recorder for MockIngredientRepository
type MockIngredientRepositoryMockRecorder struct {
	mock *MockIngredientRepository
}

// NewMockIngredientRepository creates a new mock instance
func NewMockIngredientRepository(ctrl *gomock.Controller) *MockIngredientRepository {
	mock := &MockIngredientRepository{ctrl: ctrl}
	mock.recorder = &MockIngredientRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIngredientRepository) EXPECT() *MockIngredientRepositoryMockRecorder {
	return m.recorder
}

// GetIngredients mocks base method
func (m *MockIngredientRepository) GetIngredients(ctx context.Context, min, max *uint) ([]*entity.Ingredient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngredients", ctx, min, max)
	ret0, _ := ret[0].([]*entity.Ingredient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIngredients indicates an expected call of GetIngredients
func (mr *MockIngredientRepositoryMockRecorder) GetIngredients(ctx, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngredients", reflect.TypeOf((*MockIngredientRepository)(nil).GetIngredients), ctx, min, max)
}

// GetIngredientByID mocks base method
func (m *MockIngredientRepository) GetIngredientByID(ctx context.Context, ingredientId string) (*entity.Ingredient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngredientByID", ctx, ingredientId)
	ret0, _ := ret[0].(*entity.Ingredient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIngredientByID indicates an expected call of GetIngredientByID
func (mr *MockIngredientRepositoryMockRecorder) GetIngredientByID(ctx, ingredientId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngredientByID", reflect.TypeOf((*MockIngredientRepository)(nil).GetIngredientByID), ctx, ingredientId)
}

// UpdateIngredients mocks base method
func (m *MockIngredientRepository) UpdateIngredients(ctx context.Context, ingredientId string, weight uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIngredients", ctx, ingredientId, weight)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIngredients indicates an expected call of UpdateIngredients
func (mr *MockIngredientRepositoryMockRecorder) UpdateIngredients(ctx, ingredientId, weight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIngredients", reflect.TypeOf((*MockIngredientRepository)(nil).UpdateIngredients), ctx, ingredientId, weight)
}

// MockNutritionRepository is a mock of NutritionRepository interface
type MockNutritionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockNutritionRepositoryMockRecorder
}

// MockNutritionRepositoryMockRecorder is the mock recorder for MockNutritionRepository
type MockNutritionRepositoryMockRecorder struct {
	mock *MockNutritionRepository
}

// NewMockNutritionRepository creates a new mock instance
func NewMockNutritionRepository(ctrl *gomock.Controller) *MockNutritionRepository {
	mock := &MockNutritionRepository{ctrl: ctrl}
	mock.recorder = &MockNutritionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNutritionRepository) EXPECT() *MockNutritionRepositoryMockRecorder {
	return m.recorder
}

// GetNutritions mocks base method
func (m *MockNutritionRepository) GetNutritions(ctx context.Context) ([]*entity.Nutrition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNutritions", ctx)
	ret0, _ := ret[0].([]*entity.Nutrition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNutritions indicates an expected call of GetNutritions
func (mr *MockNutritionRepositoryMockRecorder) GetNutritions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNutritions", reflect.TypeOf((*MockNutritionRepository)(nil).GetNutritions), ctx)
}

// GetNutritionByID mocks base method
func (m *MockNutritionRepository) GetNutritionByID(ctx context.Context, nutritionId string) (*entity.Nutrition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNutritionByID", ctx, nutritionId)
	ret0, _ := ret[0].(*entity.Nutrition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNutritionByID indicates an expected call of GetNutritionByID
func (mr *MockNutritionRepositoryMockRecorder) GetNutritionByID(ctx, nutritionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNutritionByID", reflect.TypeOf((*MockNutritionRepository)(nil).GetNutritionByID), ctx, nutritionId)
}
