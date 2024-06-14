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

func TestHandler_createEvents(t *testing.T) {
	type mockBehaviorEvent func(s *mock_service.MockEvent, userId int, input augventure.Event)
	type mockBehaviorSprint func(s *mock_service.MockSprint, eventId int)

	testTable := []struct {
		name                string
		inputUserId         int
		inputBody           string
		inputEvent          augventure.Event
		mockBehaviorEvent   mockBehaviorEvent
		mockBehaviorSprint  mockBehaviorSprint
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:        "OK",
			inputUserId: 1,
			inputBody:   `{"title":"title","description":"description","start_date":"now"}`,
			inputEvent: augventure.Event{
				Title:       "title",
				Description: "description",
				Start:       "now",
			},
			mockBehaviorEvent: func(s *mock_service.MockEvent, userId int, input augventure.Event) {
				s.EXPECT().Create(userId, input).Return(1, nil)
			},
			mockBehaviorSprint: func(s *mock_service.MockSprint, eventId int) {
				s.EXPECT().Create(eventId).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"eventId":1,"sprintId":1}`,
		},
		{
			name:                "Empty Fields",
			inputUserId:         1,
			inputBody:           `{"title":"title"}`,
			mockBehaviorEvent:   func(s *mock_service.MockEvent, userId int, input augventure.Event) {},
			mockBehaviorSprint:  func(s *mock_service.MockSprint, eventId int) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:        "Service Failure",
			inputUserId: 1,
			inputBody:   `{"title":"title","description":"description","start_date":"now"}`,
			inputEvent: augventure.Event{
				Title:       "title",
				Description: "description",
				Start:       "now",
			},
			mockBehaviorEvent: func(s *mock_service.MockEvent, userId int, input augventure.Event) {
				s.EXPECT().Create(userId, input).Return(0, errors.New("service failure"))
			},
			mockBehaviorSprint:  func(s *mock_service.MockSprint, eventId int) {},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
		{
			name:        "Service Failure",
			inputUserId: 1,
			inputBody:   `{"title":"title","description":"description","start_date":"now"}`,
			inputEvent: augventure.Event{
				Title:       "title",
				Description: "description",
				Start:       "now",
			},
			mockBehaviorEvent: func(s *mock_service.MockEvent, userId int, input augventure.Event) {
				s.EXPECT().Create(userId, input).Return(1, nil)
			},
			mockBehaviorSprint: func(s *mock_service.MockSprint, eventId int) {
				s.EXPECT().Create(eventId).Return(0, errors.New("service failure"))
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

			apiEvent := mock_service.NewMockEvent(c)
			testCase.mockBehaviorEvent(apiEvent, testCase.inputUserId, testCase.inputEvent)

			apiSprint := mock_service.NewMockSprint(c)
			testCase.mockBehaviorSprint(apiSprint, 1)

			services := &service.Service{Event: apiEvent, Sprint: apiSprint}
			handler := NewHandler(services)

			// Test Server
			r := gin.New()
			r.POST("/api/events", func(c *gin.Context) {
				c.Set(userCtx, testCase.inputUserId)
				handler.createEvents(c)
			})

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/events", bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
