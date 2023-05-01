package manager

import (
	"github.com/Anarr/entain/internal/model"
	"github.com/Anarr/entain/internal/repository"
	"github.com/labstack/gommon/log"
)

type (
	Manager interface {
		ProcessRequest(request model.Request) error
		CancelRequests(limit int) error
	}

	manager struct {
		repository repository.Repository
	}
)

func New(repository repository.Repository) Manager {
	return &manager{repository: repository}
}

func (m manager) ProcessRequest(request model.Request) error {
	err := m.repository.SaveRequest(request)
	if err != nil {
		return err
	}

	return m.repository.UpdateUserBalance(request)
}

func (m manager) CancelRequests(limit int) error {
	requests, err := m.repository.GetLatestRequests(limit)
	if err != nil {
		return err
	}

	log.Infof("take latest %d requests for cancellation", len(requests))

	for _, request := range requests {
		request.SwitchState() //change win state to lost state and vice versa
		if err := m.repository.CancelRequest(*request); err != nil {
			log.Error(err)
		}
	}

	return nil
}
