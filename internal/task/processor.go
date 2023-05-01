package task

import (
	"context"
	"fmt"
	"github.com/Anarr/entain/internal/manager"
	"github.com/labstack/gommon/log"
	"time"
)

const defaultTaskCountForCancellation = 10 // 1min

type (
	Processor interface {
		Start(ctx context.Context, interval time.Duration) //interva with second
	}

	processor struct {
		manager manager.Manager
	}
)

func NewProcessor(manager manager.Manager) Processor {
	return &processor{manager: manager}
}

func (p *processor) Start(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := p.manager.CancelRequests(defaultTaskCountForCancellation); err != nil {
				log.Errorf("[unable cancel the requests]: %+v", err)
			}
		case <-ctx.Done():
			fmt.Println("context cancelled")
		}
	}
}
