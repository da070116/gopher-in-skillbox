package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	gopherinskillbox "skillbox-test"
	"skillbox-test/pkg/service"
	mock_service "skillbox-test/pkg/service/mocks"
	"testing"
)

//TestHandler_addUser - test for create endpoint
func TestHandler_addUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUser, user gopherinskillbox.User)
	emptySlice := make([]int, 0)
	testUser := gopherinskillbox.User{
		Name:    "Peter",
		Age:     22,
		Friends: emptySlice,
	}
	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            gopherinskillbox.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"name": "Peter","age":22,"friends":[]}`,
			inputUser: testUser,
			mockBehavior: func(s *mock_service.MockUser, user gopherinskillbox.User) {
				s.EXPECT().CreateUser(user).Return(testUser, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"age":22,"name":"Peter"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUser := mock_service.NewMockUser(ctrl)
			testCase.mockBehavior(mockUser, testCase.inputUser)

			services := &service.Service{User: mockUser}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/add", handler.addUser)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/add", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)

		})
	}
}
