package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"skillbox/internal/domain/model"
	mock_handler "skillbox/internal/transport/http/handler/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
)

// func TestCreateUser(t *testing.T) {
// 	type mockBehavior func(s *mock_handler.MockUserService, data dto.CreateUserDTO)

// 	testTable := []struct {
// 		name                string
// 		inputBody           string
// 		inputData           dto.CreateUserDTO
// 		mockBehavior        mockBehavior
// 		expectedStatusCode  int
// 		expectedRequestBody string
// 	}{
// 		{
// 			name:      "OK",
// 			inputBody: `{"name":"ivan","age":"33"}`,
// 			inputData: dto.CreateUserDTO{
// 				Name: "ivan",
// 				Age:  "33",
// 			},
// 			mockBehavior: func(s *mock_handler.MockUserService, data dto.CreateUserDTO) {
// 				s.EXPECT().CreateUser(context.Background(), data).Return(uint64(1), nil)
// 			},
// 			expectedStatusCode: 201,

// 			expectedRequestBody: `{id:1}`,
// 		},
// 	}
// 	for _, testCases := range testTable {
// 		t.Run(testCases.name, func(t *testing.T) {
// 			c := gomock.NewController(t)
// 			defer c.Finish()

// 			l := logrus.New()
// 			engine := gin.New()
// 			svc := mock_handler.NewMockUserService(c)
// 			testCases.mockBehavior(svc, testCases.inputData)

// 			handler := New(svc, engine, l)
// 			engine.POST("/create", handler.CreateUser)
// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest("POST", "/create", bytes.NewBufferString(testCases.inputBody))
// 			engine.ServeHTTP(w, req)

// 			assert.Equal(t, w.Code, testCases.expectedStatusCode)
// 			assert.Equal(t, w.Body, testCases.expectedRequestBody)

// 		})
// 	}
// }

func TestCreateUser(t *testing.T) {
	log := logrus.New()
	eng := gin.Default()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	exp := uint64(1)
	data := model.User{
		Age:  33,
		Name: "stas",
	}
	svc := mock_handler.NewMockUserService(ctrl)
	svc.EXPECT().CreateUser(ctx, data).Return(exp, nil)
	b, err := json.Marshal(data)
	if err != nil {
		t.Error()
	}
	h := New(svc, eng, log)
	eng.POST("/create", h.CreateUser)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/create", bytes.NewBuffer(b))

	eng.ServeHTTP(rec, req)
	responce := rec.Result()
	defer responce.Body.Close()
	dd, _ := ioutil.ReadAll(responce.Body)
	assert.Equal(t, string(dd), exp)
	//assert.Equal(t, rec.Code, http.StatusCreated)
}
