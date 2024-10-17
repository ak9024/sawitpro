package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ak9024/sawitpro/generated"
	"github.com/ak9024/sawitpro/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/mock/gomock"
)

var (
	server         *Server
	mockRepository *repository.MockRepositoryInterface
)

func beforeEach(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository = repository.NewMockRepositoryInterface(ctrl)
	server = NewServer(NewServerOptions{
		Repository: mockRepository,
	})

	return func() {}
}

func TestCreateEstate(t *testing.T) {
	t.Run("Should be create a new estate", func(t *testing.T) {
		Convey("POST /estate", t, func() {
			type (
				args struct {
					payload string
				}
			)

			testCases := []struct {
				testiD         int
				testDesc       string
				args           args
				mockFunc       func()
				wantStatusCode int
				wantResp       generated.EstateResponse
			}{
				{
					testiD:   1,
					testDesc: "Success - 200 ok!",
					args: args{
						payload: `{  "width": 10, "length": 1 }`,
					},
					mockFunc: func() {
						mockRepository.EXPECT().CreateEstate(gomock.Any(), gomock.Any()).Return(repository.Estate{}, nil)
					},
					wantStatusCode: http.StatusOK,
				},
			}

			for _, tc := range testCases {
				before := beforeEach(t)
				defer before()

				Convey(fmt.Sprintf("%d: %s", tc.testiD, tc.testDesc), func() {
					tc.mockFunc()

					method := echo.POST
					path := "/estate"

					e := echo.New()
					req := httptest.NewRequest(method, path, bytes.NewReader([]byte(tc.args.payload)))
					req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
					rr := httptest.NewRecorder()
					c := e.NewContext(req, rr)
					_ = server.PostEstate(c)

					var resp generated.EstateResponse
					_ = json.Unmarshal(rr.Body.Bytes(), &resp)
					t.Log(tc.wantResp)
					So(rr.Code, ShouldEqual, tc.wantStatusCode)
				})
			}
		})
	})
}

func TestCreateTree(t *testing.T) {
	t.Run("Should be create a new tree", func(t *testing.T) {
		Convey("POST /estate/<id>/tree", t, func() {
			type (
				args struct {
					payload string
				}
			)

			testCases := []struct {
				testiD         int
				testDesc       string
				args           args
				mockFunc       func()
				wantStatusCode int
				wantResp       generated.EstateTreeResponse
			}{
				{
					testiD:   1,
					testDesc: "Success - 200 ok!",
					args: args{
						payload: `{  "x": 10, "y": 1, "height": 30 }`,
					},
					mockFunc: func() {
						mockRepository.EXPECT().CreateEstateTree(gomock.Any(), gomock.Any()).Return(repository.EstateTree{}, nil)
					},
					wantStatusCode: http.StatusOK,
					wantResp: generated.EstateTreeResponse{
						Id: "uuid",
					},
				},
			}

			for _, tc := range testCases {
				before := beforeEach(t)
				defer before()

				Convey(fmt.Sprintf("%d: %s", tc.testiD, tc.testDesc), func() {
					tc.mockFunc()

					method := echo.POST
					id := "uuid"
					path := fmt.Sprintf("/estate/%s/tree", id)

					e := echo.New()
					req := httptest.NewRequest(method, path, bytes.NewReader([]byte(tc.args.payload)))
					req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
					rr := httptest.NewRecorder()
					c := e.NewContext(req, rr)
					_ = server.PostEstateIdTree(c, id)

					var resp generated.EstateTreeResponse
					_ = json.Unmarshal(rr.Body.Bytes(), &resp)
					So(rr.Code, ShouldEqual, tc.wantStatusCode)
					So(resp.Id, ShouldEqual, tc.wantResp.Id)
				})
			}
		})
	})
}

func TestGetStat(t *testing.T) {
	t.Run("Should be get a stats from estate", func(t *testing.T) {
		Convey("GET /estate/id/stats", t, func() {

			testCases := []struct {
				testiD         int
				testDesc       string
				mockFunc       func()
				wantStatusCode int
				wantResp       generated.EstateStatsResponse
			}{
				{
					testiD:   1,
					testDesc: "Success - 200 ok!",
					mockFunc: func() {
						mockRepository.EXPECT().GetStats(gomock.Any(), gomock.Any()).Return(0, 0, 0, 0.5, nil)
					},
					wantStatusCode: http.StatusOK,
					wantResp: generated.EstateStatsResponse{
						Count:  0,
						Max:    0,
						Min:    0,
						Median: 0,
					},
				},
			}

			for _, tc := range testCases {
				before := beforeEach(t)
				defer before()

				Convey(fmt.Sprintf("%d: %s", tc.testiD, tc.testDesc), func() {
					tc.mockFunc()

					method := echo.GET
					id := uuid.New().String()
					path := fmt.Sprintf("/estate/%s/stats", id)

					e := echo.New()
					req := httptest.NewRequest(method, path, nil)
					req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
					rr := httptest.NewRecorder()
					c := e.NewContext(req, rr)
					_ = server.GetEstateIdStats(c, id)

					var resp generated.EstateStatsResponse
					_ = json.Unmarshal(rr.Body.Bytes(), &resp)
					So(rr.Code, ShouldEqual, tc.wantStatusCode)
					So(resp, ShouldEqual, tc.wantResp)
				})
			}
		})
	})
}
