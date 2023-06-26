package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"skillbox/internal/transport/http/dto"
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
	//init dependences
	log := logrus.New()
	eng := gin.Default()
	ctx := context.Background()
	exp := uint64(1)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//create mock for service method
	data := dto.CreateUserDTO{
		Age:  "33",
		Name: "stas",
	}
	svc := mock_handler.NewMockUserService(ctrl)
	svc.EXPECT().CreateUser(ctx, data).Return(exp, nil)

	//create handler
	h := New(svc, eng, log)
	//register handle func
	eng.POST("/create", h.CreateUser)
	//create request data
	reqdata := dto.CreateUserDTO{
		Age:  "33",
		Name: "stas",
	}
	fmt.Println("reqdata:", reqdata)
	reqbytes, err := json.Marshal(reqdata)
	if err != nil {
		t.Error()
	}
	fmt.Println("reqbytes:", reqbytes)
	//doing request
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/create", bytes.NewBuffer(reqbytes))
	eng.ServeHTTP(rec, req)
	//read responce
	responce := rec.Result()
	defer responce.Body.Close()
	dd, _ := ioutil.ReadAll(responce.Body)
	//check result
	assert.Equal(t, string(dd), exp)
	//assert.Equal(t, rec.Code, http.StatusCreated)
}
