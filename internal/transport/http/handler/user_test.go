package handler

import (
	"bytes"
	"context"
	"net/http/httptest"
	"skillbox/internal/transport/http/dto"
	mock_handler "skillbox/internal/transport/http/handler/mocks"
	"skillbox/internal/transport/http/middleware"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
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
			inputBody: `{"name":"ivan","age":"33"}`,
			inputData: dto.CreateUserDTO{
				Name: "ivan",
				Age:  "33",
			},
			mockBehavior: func(s *mock_handler.MockUserService, data dto.CreateUserDTO) {
				s.EXPECT().CreateUser(context.Background(), data).Return(uint64(1), nil)
			},
			expectedStatusCode: 201,

			expectedRequestBody: `{id:1}`,
		},
	}
	for _, testCases := range testTable {
		t.Run(testCases.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			l := logrus.New()
			middle := middleware.New(l)
			engine := gin.Default()
			svc := mock_handler.NewMockUserService(c)
			testCases.mockBehavior(svc, testCases.inputData)

			handler := New(svc, engine, middle, l)
			engine.POST("/create", handler.CreateUser)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/create", bytes.NewBufferString(testCases.inputBody))
			engine.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCases.expectedStatusCode)
			assert.Equal(t, w.Body, testCases.expectedRequestBody)

		})
	}
}
