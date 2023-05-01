package model

import "time"

type (
	User struct {
		ID        int       `json:"id" gorm:"primaryKey"`
		Balance   float64   `json:"balance"`
		CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	}

	Request struct {
		ID            int       `json:"id" gorm:"primaryKey"`
		UserID        int       `json:"user_id"`
		TransactionID string    `json:"transaction_id" gorm:"unique"`
		State         string    `json:"state"`
		Amount        float64   `json:"amount"`
		Processed     bool      `json:"processed"`
		CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	}

	APIError struct {
		StatusCode int
		Err        string
	}
)

func (r *Request) SwitchState() {
	if r.State == "win" {
		r.State = "lost"
	} else {
		r.State = "win"
	}
}

func (ae APIError) Error() string {
	return ae.Err
}
