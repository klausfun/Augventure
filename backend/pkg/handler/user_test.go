package handler

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	augventure "github.com/klausfun/Augventure"
	"github.com/klausfun/Augventure/pkg/service"
	mock_service "github.com/klausfun/Augventure/pkg/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_getUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockProfile, userId int)

	testTable := []struct {
		name                string
		inputUserId         int
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:        "OK",
			inputUserId: 1,
			mockBehavior: func(s *mock_service.MockProfile, userId int) {
				s.EXPECT().GetById(userId).Return(augventure.Author{
					Name:     "",
					Username: "test",
					Email:    "test@mail.ru",
					PfpUrl:   "",
					Bio:      "",
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"user":{"name":"","username":"test","email":"test@mail.ru","pfp_url":"","bio":""}}`,
		},
		{
			name:        "Service Failure",
			inputUserId: 1,
			mockBehavior: func(s *mock_service.MockProfile, userId int) {
				s.EXPECT().GetById(userId).Return(augventure.Author{
					Name:     "",
					Username: "test",
					Email:    "test@mail.ru",
					PfpUrl:   "",
					Bio:      "",
				}, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			apiUser := mock_service.NewMockProfile(c)
			testCase.mockBehavior(apiUser, testCase.inputUserId)

			services := &service.Service{Profile: apiUser}
			handler := NewHandler(services)

			// Test Server
			r := gin.New()
			r.GET("/api/users/me", func(c *gin.Context) {
				c.Set(userCtx, testCase.inputUserId)
				handler.getUser(c)
			})

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/users/me", bytes.NewBufferString(""))

			// Perform Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
