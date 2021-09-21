package server

import (
	"FPproject/Backend/models"
	"FPproject/Backend/server/mock"
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/gin-gonic/gin"
)

type test struct {
	name           string
	body           string
	mock           func(*mock.Mockrepository)
	want           string
	wantStatusCode int
}

func TestInsertUser(t *testing.T) {
	tests := []test{
		{
			"success case",
			`{
				"username": "user2",
				"name": "nameofuser2",
				"password": "cde123",
				"type": "customer"
			}`,
			func(m *mock.Mockrepository) {
				m.EXPECT().InsertUser(models.User{
					Username: "user2",
					Name:     "nameofuser2",
					Password: "cde123",
					UserType: "customer",
				}).Return("randomstring", nil)
			}, `{"id":"randomstring","status":"OK"}`, http.StatusOK,
		},
		// {
		// 	"failure case - dependency",
		// 	`{
		// 		"username": "user2",
		// 		"name": "nameofuser2",
		// 		"password": "cde123",
		// 		"type": "customer"
		// 	}`,
		// 	func(m *mock.Mockrepository) {
		// 		m.EXPECT().InsertUser(models.User{
		// 			Username: "user2",
		// 			Name:     "nameofuser2",
		// 			Password: "cde123",
		// 			UserType: "customer",
		// 		}).Return("", errors.New("unknown error"))
		// 	}, `{"status":"internal server error"}`, http.StatusInternalServerError,
		// },
		// {
		// 	"failure case - dependency",
		// 	`{}`,
		// 	func(m *mock.Mockrepository) {},
		// 	`{"status":"bad request"}`,
		// 	http.StatusBadRequest,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockrepository(ctrl)
			tt.mock(m)
			h := handler(m)
			// create a mock http object
			res := httptest.NewRecorder()
			b := bytes.NewBufferString(tt.body)
			req, err := http.NewRequest(http.MethodPost, "localhost:8080/user", b)
			if err != nil {
				panic(err)
			}
			_, r := gin.CreateTestContext(res)
			r.POST("/user", h.InsertUser)
			r.ServeHTTP(res, req)
			if strings.Compare(tt.want, res.Body.String()) != 0 {
				t.Errorf("want: %s, got: %s", tt.want, res.Body.String())
			}
			if tt.wantStatusCode != res.Code {
				t.Errorf("want: %d, got: %d", tt.wantStatusCode, res.Code)
			}
		})
	}
}
