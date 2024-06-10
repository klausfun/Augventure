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

func TestHandler_getSuggestionsBySprintId(t *testing.T) {
	type mockBehavior func(s *mock_service.MockSuggestion, sprintId int)

	testTable := []struct {
		name                string
		inputBody           string
		inputSprintId       int
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:          "OK",
			inputBody:     `{"sprint_id":1}`,
			inputSprintId: 1,
			mockBehavior: func(s *mock_service.MockSuggestion, sprintId int) {
				s.EXPECT().GetBySprintId(sprintId).Return([]augventure.FilterSuggestions{
					{
						Id:       1,
						AuthorId: 1,
						SprintId: 1,
						Author: augventure.Author{
							Name:     "",
							Username: "test",
							Email:    "test@mail.ru",
							PfpUrl:   "",
							Bio:      "",
						},
						Content:  "content",
						PostDate: "",
						Votes:    0,
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"data":[{"id":1,"author_id":1,"sprint_id":1,"author":{"name":"","username":"test","email":"test@mail.ru","pfp_url":"","bio":""},"content":"content","post_date":"","votes":0}]}`,
		},
		{
			name:          "Service Failure",
			inputBody:     `{"sprint_id":1}`,
			inputSprintId: 1,
			mockBehavior: func(s *mock_service.MockSuggestion, sprintId int) {
				s.EXPECT().GetBySprintId(sprintId).Return([]augventure.FilterSuggestions{
					{
						Id:       1,
						AuthorId: 1,
						SprintId: 1,
						Author: augventure.Author{
							Name:     "",
							Username: "test",
							Email:    "test@mail.ru",
							PfpUrl:   "",
							Bio:      "",
						},
						Content:  "content",
						PostDate: "",
						Votes:    0,
					},
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

			apiSuggestion := mock_service.NewMockSuggestion(c)
			testCase.mockBehavior(apiSuggestion, testCase.inputSprintId)

			services := &service.Service{Suggestion: apiSuggestion}
			handler := NewHandler(services)

			// Test Server
			r := gin.New()
			r.POST("/api/suggestions/get", func(c *gin.Context) {
				c.Set(userCtx, testCase.inputSprintId)
				handler.getSuggestionsBySprintId(c)
			})

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/suggestions/get", bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
