package integration_tests

import (
	"fmt"
	mocks "fullstack/backend/mocks"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

func Test_registerUser(t *testing.T) {
	type mockBehaviour func(mus *mocks.MockUserService)

	tests := []struct {
		name          string
		expRespCode   int
		expRespBody   string
		reqBody       string
		mockBehaviour mockBehaviour
	}{
		{
			name:        "All fields: Email, Password",
			expRespCode: http.StatusCreated,
			reqBody: fmt.Sprintf(
				`{"Email":"%s", "Password":"%s"}`,
				"test1@test.com", "12345678",
			),
			expRespBody: `{"status":"success","data":{"message":"successful registration"}}`,
		},
		{
			name:        "Without Email",
			expRespCode: http.StatusBadRequest,
			reqBody: fmt.Sprintf(
				`{"Password":"%s"}`,
				"12345678",
			),
			expRespBody: `{"status":"error","message":"invalid email"}`,
		},
		{
			name:        "Without Password",
			expRespCode: http.StatusBadRequest,
			reqBody: fmt.Sprintf(
				`{"Email":"%s"}`,
				"test2@test.com",
			),
			expRespBody: `{"status":"error","message":"invalid password"}`,
		},
		{
			name:        "Duplicate Email",
			expRespCode: http.StatusBadRequest,
			reqBody: fmt.Sprintf(
				`{"Email":"%s", "Password":"%s"}`,
				"test1@gmail.com", "12345678",
			),
			expRespBody: `{"status":"error","message":"UNIQUE constraint failed: users.email"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			initTests(t)

			url := "http://" + testURL + "/register"
			method := "POST"

			req, err := http.NewRequest(method, url, strings.NewReader(test.reqBody))
			assert.NoError(t, err)

			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
				return
			}
			assert.NoError(t, err)

			bytes, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, test.expRespCode, res.StatusCode)
			assert.JSONEq(t, test.expRespBody, string(bytes))
		})
	}

}

//func Test_loginUser(t *testing.T) {
//	type mockBehaviour func(mus *mocks.MockUserService, mam *mocks.MockAuthManager, expAuthzToken string)
//
//	tests := []struct {
//		name               string
//		expRespCode        int
//		expRespBody        string
//		expRespAuthzHeader string
//		expAuthzToken      string
//		reqBody            string
//		mockBehaviour      mockBehaviour
//	}{
//		{
//			name:        "All fields: Email, Password",
//			expRespCode: http.StatusOK,
//			reqBody: fmt.Sprintf(
//				`{"Email":"%s", "Password":"%s"}`,
//				mockUserLogin1.Email, mockUserLogin1.Password,
//			),
//			expRespBody:        `{"status":"success","data":{"message":"successful login"}}`,
//			expRespAuthzHeader: "Bearer correct_token",
//			expAuthzToken:      "correct_token",
//			mockBehaviour: func(mus *mocks.MockUserService, mam *mocks.MockAuthManager, expAuthzToken string) {
//				mus.On("Login", &mockUserLogin1).Return(mockUser1.ID, nil)
//				mam.On("MakeAuthn", mockUser1.ID).Return(expAuthzToken, nil)
//			},
//		},
//		{
//			name:        "Without Email",
//			expRespCode: http.StatusBadRequest,
//			reqBody: fmt.Sprintf(
//				`{"Password":"%s"}`,
//				mockUserReg2.Password,
//			),
//			expRespBody:        `{"status":"error","message":"invalid email"}`,
//			expRespAuthzHeader: "",
//			expAuthzToken:      "",
//			mockBehaviour: func(mus *mocks.MockUserService, mam *mocks.MockAuthManager, expAuthzToken string) {
//				mus.On("Login", &mockUserLogin2).Return(uint(0), entity.ErrInvalidEmail)
//			},
//		},
//		{
//			name:        "Without Password",
//			expRespCode: http.StatusBadRequest,
//			reqBody: fmt.Sprintf(
//				`{"Email":"%s"}`,
//				mockUserLogin3.Email,
//			),
//			expRespBody:        `{"status":"error","message":"invalid password"}`,
//			expRespAuthzHeader: "",
//			expAuthzToken:      "",
//			mockBehaviour: func(mus *mocks.MockUserService, mam *mocks.MockAuthManager, expAuthzToken string) {
//				mus.On("Login", &mockUserLogin3).Return(uint(0), entity.ErrInvalidPassword)
//			},
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			mockUserService := new(mocks.MockUserService)
//			mockVisitService := new(mocks.MockVisitService)
//			mockAuthManager := new(mocks.MockAuthManager)
//			mockOAuthManager := new(mocks.MockOAuthManager)
//
//			test.mockBehaviour(mockUserService, mockAuthManager, test.expAuthzToken)
//
//			handler := &Handler{mockUserService, mockVisitService, mockAuthManager, mockOAuthManager}
//
//			r := mux.NewRouter()
//			r.HandleFunc("/user/login", handler.loginUser).Methods(http.MethodPost)
//
//			w := httptest.NewRecorder()
//			req := httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader(test.reqBody))
//			r.ServeHTTP(w, req)
//
//			assert.Equal(t, test.expRespCode, w.Code)
//			assert.Equal(t, test.expRespAuthzHeader, w.Header().Get("Authorization"))
//			assert.JSONEq(t, test.expRespBody, w.Body.String())
//		})
//	}
//
//}
