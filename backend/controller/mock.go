// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package controller is a generated GoMock package.
package controller

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "ref/entity"
	reflect "reflect"
)

// MockIngredientUseCase is a mock of IngredientUseCase interface
type MockIngredientUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIngredientUseCaseMockRecorder
}

// MockIngredientUseCaseMockRecorder is the mock recorder for MockIngredientUseCase
type MockIngredientUseCaseMockRecorder struct {
	mock *MockIngredientUseCase
}

// NewMockIngredientUseCase creates a new mock instance
func NewMockIngredientUseCase(ctrl *gomock.Controller) *MockIngredientUseCase {
	mock := &MockIngredientUseCase{ctrl: ctrl}
	mock.recorder = &MockIngredientUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIngredientUseCase) EXPECT() *MockIngredientUseCaseMockRecorder {
	return m.recorder
}

// GetIngredients mocks base method
func (m *MockIngredientUseCase) GetIngredients(ctx context.Context, min, max *uint) ([]*entity.Ingredient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngredients", ctx, min, max)
	ret0, _ := ret[0].([]*entity.Ingredient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIngredients indicates an expected call of GetIngredients
func (mr *MockIngredientUseCaseMockRecorder) GetIngredients(ctx, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngredients", reflect.TypeOf((*MockIngredientUseCase)(nil).GetIngredients), ctx, min, max)
}

// UpdateIngredients mocks base method
func (m *MockIngredientUseCase) UpdateIngredients(ctx context.Context, ingredientId string, weight uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIngredients", ctx, ingredientId, weight)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIngredients indicates an expected call of UpdateIngredients
func (mr *MockIngredientUseCaseMockRecorder) UpdateIngredients(ctx, ingredientId, weight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIngredients", reflect.TypeOf((*MockIngredientUseCase)(nil).UpdateIngredients), ctx, ingredientId, weight)
}

// MockNutritionUseCase is a mock of NutritionUseCase interface
type MockNutritionUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockNutritionUseCaseMockRecorder
}

// MockNutritionUseCaseMockRecorder is the mock recorder for MockNutritionUseCase
type MockNutritionUseCaseMockRecorder struct {
	mock *MockNutritionUseCase
}

// NewMockNutritionUseCase creates a new mock instance
func NewMockNutritionUseCase(ctrl *gomock.Controller) *MockNutritionUseCase {
	mock := &MockNutritionUseCase{ctrl: ctrl}
	mock.recorder = &MockNutritionUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNutritionUseCase) EXPECT() *MockNutritionUseCaseMockRecorder {
	return m.recorder
}

// GetNutritions mocks base method
func (m *MockNutritionUseCase) GetNutritions(ctx context.Context) ([]*entity.Nutrition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNutritions", ctx)
	ret0, _ := ret[0].([]*entity.Nutrition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNutritions indicates an expected call of GetNutritions
func (mr *MockNutritionUseCaseMockRecorder) GetNutritions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNutritions", reflect.TypeOf((*MockNutritionUseCase)(nil).GetNutritions), ctx)
}

// GetNutritionByID mocks base method
func (m *MockNutritionUseCase) GetNutritionByID(ctx context.Context, nutritionId string) (*entity.Nutrition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNutritionByID", ctx, nutritionId)
	ret0, _ := ret[0].(*entity.Nutrition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNutritionByID indicates an expected call of GetNutritionByID
func (mr *MockNutritionUseCaseMockRecorder) GetNutritionByID(ctx, nutritionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNutritionByID", reflect.TypeOf((*MockNutritionUseCase)(nil).GetNutritionByID), ctx, nutritionId)
}