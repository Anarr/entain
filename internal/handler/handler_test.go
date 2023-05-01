package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/Anarr/entain/internal/handler"
	"github.com/Anarr/entain/internal/model"
	mockedmanager "github.com/Anarr/entain/internal/test/mocks/manager"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_ProcessRequestHandler(t *testing.T) {
	testCases := []struct {
		name       string
		input      *handler.ProcessRequest
		header     string
		statusCode int
	}{
		{
			name: "TEST success",
			input: &handler.ProcessRequest{
				State:         "win",
				TransactionID: "tr-1",
				Amount:        10,
			},
			header:     "game",
			statusCode: http.StatusOK,
		},
		{
			name: "TEST wrong amount",
			input: &handler.ProcessRequest{
				State:         "hey",
				TransactionID: "tr-1",
				Amount:        0,
			},
			header:     "server",
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "TEST without header",
			input: &handler.ProcessRequest{
				State:         "hey",
				TransactionID: "tr-1",
				Amount:        10,
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "TEST invalid header",
			input: &handler.ProcessRequest{
				State:         "hey",
				TransactionID: "tr-1",
				Amount:        0,
			},
			header:     "invalid",
			statusCode: http.StatusUnprocessableEntity,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			bs, err := json.Marshal(tc.input)
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/process", bytes.NewBuffer(bs))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set("Source-Type", tc.header)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			managerMock := mockedmanager.NewManager(t)
			if tc.statusCode == http.StatusOK {
				rm := model.Request{
					TransactionID: tc.input.TransactionID,
					State:         tc.input.State,
					Amount:        tc.input.Amount,
					UserID:        1,
				}
				managerMock.On("ProcessRequest", rm).Return(nil)
			}

			h := handler.New(managerMock)
			err = h.ProcessRequestHandler(c)

			if tc.statusCode != http.StatusOK {
				assert.Error(t, err)
				he, ok := err.(model.APIError)
				assert.True(t, ok)
				assert.Equal(t, tc.statusCode, he.StatusCode)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.statusCode, rec.Code)
			}
		})
	}
}
