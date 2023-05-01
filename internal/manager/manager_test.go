package manager_test

import (
	"errors"
	"github.com/Anarr/entain/internal/manager"
	"github.com/Anarr/entain/internal/model"
	mockedrepository "github.com/Anarr/entain/internal/test/mocks/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManager_ProcessRequest(t *testing.T) {

	testCases := []struct {
		name    string
		request model.Request
		err     error
	}{
		{
			name: "Test Process Request successfully",
			request: model.Request{
				UserID:        1,
				TransactionID: "tr-1",
				Amount:        10,
			},
			err: nil,
		},
		{
			name: "Test Process Request error",
			request: model.Request{
				UserID:        0,
				TransactionID: "tr-2",
				Amount:        10,
			},
			err: errors.New("some error"),
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			repositoryMock := mockedrepository.NewRepository(t)
			if tc.err == nil {
				repositoryMock.On("UpdateUserBalance", tc.request).Return(tc.err)
			}
			repositoryMock.On("SaveRequest", tc.request).Return(tc.err)

			m := manager.New(repositoryMock)
			err := m.ProcessRequest(tc.request)

			if tc.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestManager_CancelRequests(t *testing.T) {
	repositoryMock := mockedrepository.NewRepository(t)
	repositoryMock.On("GetLatestRequests", 1).Return(nil, nil)

	m := manager.New(repositoryMock)
	err := m.CancelRequests(1)
	assert.NoError(t, err)
}
