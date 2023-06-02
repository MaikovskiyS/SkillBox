package handler

import (
	"context"
	"skillbox/internal/transport/http/dto"
	mock_handler "skillbox/internal/transport/http/handler/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	type mockBehavior func(s *mock_handler.MockUserService, data dto.CreateUserDTO)

	testTable := []struct {
		name                string
		inputBody           string
		inputData           dto.CreateUserDTO
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"stas","age":"29"}`,
			inputData: dto.CreateUserDTO{
				Name: "stas",
				Age:  "29",
			},
			mockBehavior: func(s *mock_handler.MockUserService, data dto.CreateUserDTO) {
				s.EXPECT().CreateUser(context.Background(), data).Return(1, nil)
			},
			expectedStatusCode: 201,

			expectedRequestBody: `{id:1}`,
		},
	}
	for _, testCases := range testTable {
		t.Run(testCases.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			svc := mock_handler.NewMockUserService(c)
			testCases.mockBehavior(svc, testCases.inputData)
		})
	}
}
