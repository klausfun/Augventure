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

func TestHandler_updatePassword(t *testing.T) {
	type mockBehavior func(s *mock_service.MockProfile, userId int, passwords augventure.UpdatePasswordInput)

	testTable := []struct {
		name                string
		inputBody           string
		inputUserId         int
		inputPasswords      augventure.UpdatePasswordInput
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:        "OK",
			inputUserId: 1,
			inputBody:   `{"old_password":"old_password", "new_password":"new_password"}`,
			inputPasswords: augventure.UpdatePasswordInput{
				OldPassword: "old_password",
				NewPassword: "new_password",
			},
			mockBehavior: func(s *mock_service.MockProfile, userId int, passwords augventure.UpdatePasswordInput) {
				s.EXPECT().UpdatePassword(userId, passwords).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"status":"ok"}`,
		},
		{
			name:                "Empty Fields",
			inputUserId:         1,
			inputBody:           `{"old_password":"old_password"`,
			mockBehavior:        func(s *mock_service.MockProfile, userId int, passwords augventure.UpdatePasswordInput) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:        "Service Failure",
			inputUserId: 1,
			inputBody:   `{"old_password":"old_password", "new_password":"new_password"}`,
			inputPasswords: augventure.UpdatePasswordInput{
				OldPassword: "old_password",
				NewPassword: "new_password",
			},
			mockBehavior: func(s *mock_service.MockProfile, userId int, passwords augventure.UpdatePasswordInput) {
				s.EXPECT().UpdatePassword(userId, passwords).Return(errors.New("service failure"))
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
			testCase.mockBehavior(apiUser, testCase.inputUserId, testCase.inputPasswords)

			services := &service.Service{Profile: apiUser}
			handler := NewHandler(services)

			// Test Server
			r := gin.New()
			r.PUT("/api/users/me/password_reset", func(c *gin.Context) {
				c.Set(userCtx, testCase.inputUserId)
				handler.updatePassword(c)
			})

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/api/users/me/password_reset",
				bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
