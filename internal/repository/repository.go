package repository

import (
	"errors"
	"github.com/Anarr/entain/internal/model"
	"gorm.io/gorm"
)

type (
	Repository interface {
		SaveRequest(m model.Request) error
		UpdateUserBalance(m model.Request) error
		CancelRequest(m model.Request) error
		GetLatestRequests(limit int) ([]*model.Request, error)
	}

	repository struct {
		db *gorm.DB
	}
)

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) SaveRequest(m model.Request) error {
	res := r.db.Save(&m)
	return res.Error
}

func (r *repository) UpdateUserBalance(m model.Request) error {
	if m.State == "win" {
		return r.db.Model(&model.User{ID: m.UserID}).
			Update("balance", gorm.Expr("balance + ?", m.Amount)).
			Error
	}

	user, err := r.getUser(m.UserID)
	if err != nil {
		return err
	}

	if user.Balance < m.Amount {
		return errors.New("lost amount can not be great user balance")
	}

	return r.db.Model(&model.User{ID: m.UserID}).
		Update("balance", gorm.Expr("balance - ?", m.Amount)).
		Error
}

func (r *repository) CancelRequest(m model.Request) error {
	tx := r.db.Begin()
	defer tx.Rollback()

	if err := r.UpdateUserBalance(m); err != nil {
		return err
	}

	res := tx.Model(&model.Request{ID: m.ID}).Update("processed", true)

	if res.Error != nil {
		return res.Error
	}

	tx.Commit()

	return nil
}

func (r *repository) GetLatestRequests(limit int) ([]*model.Request, error) {
	var requests []*model.Request
	res := r.db.Limit(limit).Where("id%2 = 1 AND processed = false").Find(&requests).Order("id desc")
	if res.Error != nil {
		return nil, res.Error
	}

	return requests, nil
}

func (r *repository) getUser(userID int) (*model.User, error) {
	var u model.User
	res := r.db.Find(&u, &model.User{ID: userID})
	if res.Error != nil {
		return nil, res.Error
	}

	return &u, nil
}
