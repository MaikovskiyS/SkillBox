package createHandler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"skillbox/internal/domain/model"
	mock_service "skillbox/internal/domain/service/mocks"
	"skillbox/internal/transport/http/handler"
	createHandler "skillbox/internal/transport/http/handler/createUser"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
)

func TestCreate(t *testing.T) {
	type mockBihavior func(svc *mock_service.MockUserService, svcin *model.User, svcout uint64)

	tables := []struct {
		name     string
		dto      createHandler.DTO
		svcin    *model.User
		svcout   uint64
		bihavior mockBihavior
		expcode  int
		exp      string
	}{{
		name: "ok",
		dto: createHandler.DTO{
			Name: "ivan",
			Age:  "23",
		},
		svcin: &model.User{
			Name: "ivan",
			Age:  23,
		},
		svcout: 1,
		bihavior: func(svc *mock_service.MockUserService, in *model.User, out uint64) {

			svc.EXPECT().CreateUser(gomock.Any(), in).Return(out, nil)
		},
		expcode: http.StatusCreated,
		exp:     `User created. ID:1`,
	}, {
		name: "validation err",
		dto: createHandler.DTO{
			Name: "123",
			Age:  "asd",
		},
		expcode: http.StatusBadRequest,
		exp:     `create User validation err:%!(EXTRA string=age: must contain digits only; name: must contain English letters only.)`,
	}}
	//check usecases
	for _, usecase := range tables {
		t.Run(usecase.name, func(t *testing.T) {
			//init dependences
			l := logrus.New()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			//create mock for service method
			svc := mock_service.NewMockUserService(ctrl)
			if usecase.name == "ok" {
				fmt.Println("first ok test")
				usecase.bihavior(svc, usecase.svcin, usecase.svcout)
			}
			//create router and register handle func
			router := handler.New(svc, l)
			router.Engine.POST("/user", router.Creator.User)

			//doing request
			reqbytes, err := json.Marshal(usecase.dto)
			if err != nil {
				t.Error()
			}
			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(reqbytes))
			router.Engine.ServeHTTP(rec, req)
			//read responce
			responce := rec.Result()
			defer responce.Body.Close()
			body, _ := ioutil.ReadAll(responce.Body)
			//fmt.Println("responce:", string(body))
			//check result
			fmt.Println(body)
			assert.Equal(t, responce.StatusCode, usecase.expcode)
			//assert.Equal(t, string(body), usecase.exp)
		})
	}
}
