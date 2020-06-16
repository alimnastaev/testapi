package router_test

import (
	"github.com/alimnastaev/testapi/internal/api/router/_mocks/generated/controller"
	"github.com/golang/mock/gomock"

	"github.com/alimnastaev/testapi/internal/api/router"
	"github.com/stretchr/testify/assert"

	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Healthcheck_Normally_WritesBackA200Status(t *testing.T) {
	r := router.New(nil, "1.0.0")

	mockResp := httptest.NewRecorder()

	r.Healthcheck(mockResp, nil)

	body, _ := ioutil.ReadAll(mockResp.Body)
	assert.Equal(t, "Here is the version: 1.0.0", string(body))
	assert.Equal(t, http.StatusOK, mockResp.Code)
}

func Test_Login_OnErrorFromController_ReturnsError(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockC := mock_controller.NewMockAPI(mockController)
	mockC.EXPECT().Login(gomock.Any(), gomock.Any()).Times(1).Return("", errors.New("test-error"))

	mockResp := httptest.NewRecorder()

	req := router.LoginRequest{}

	body, _ := json.Marshal(req)

	request := http.Request{Body: ioutil.NopCloser(bytes.NewReader(body))}

	r := router.New(mockC, "1.0.0")
	r.Login(mockResp, &request)

	assert.Equal(t, http.StatusInternalServerError, mockResp.Code)
}

func Test_Login_Normally_ReturnsError(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockC := mock_controller.NewMockAPI(mockController)
	mockC.EXPECT().Login("email", "pass").Times(1).Return("token", nil)

	mockResp := httptest.NewRecorder()

	req := router.LoginRequest{
		Email:    "email",
		Password: "pass",
	}

	body, _ := json.Marshal(req)

	request := http.Request{Body: ioutil.NopCloser(bytes.NewReader(body))}

	r := router.New(mockC, "1.0.0")
	r.Login(mockResp, &request)

	var resp router.LoginResponse
	err := json.NewDecoder(mockResp.Body).Decode(&resp)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, mockResp.Code)
	assert.Equal(t, "token", resp.AuthToken)
}
