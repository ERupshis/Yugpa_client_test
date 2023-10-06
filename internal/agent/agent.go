package agent

import (
	"context"
	"time"

	"github.com/erupshis/yugpa_test/internal/agent/dialer"
	"github.com/erupshis/yugpa_test/internal/agent/taskgenerator"
	"github.com/erupshis/yugpa_test/internal/logger"
	"github.com/erupshis/yugpa_test/internal/messages"
)

type Agent struct {
	dialer dialer.BaseDialer

	log logger.BaseLogger
}

func Create(baseDialer dialer.BaseDialer, baseLogger logger.BaseLogger) Agent {
	return Agent{
		dialer: baseDialer,
		log:    baseLogger,
	}
}

func (a *Agent) Run(ctx context.Context, connCount int) {
	a.log.Info("[Agent:Run] starting '%d' goroutines", connCount)
	for i := 0; i < connCount; i++ {
		go func(num int) {
			for {
				select {
				case <-ctx.Done():
					a.log.Info("[Agent:Run] goroutine '%d' has been stopped outside by context cancellation")
					return
				default:
					reqMsg := messages.Request{}
					reqMsg.Path = taskgenerator.GenerateRandomPath()
					resp, err := a.dialer.MakeRequestToServer(ctx, &reqMsg)
					if err != nil {
						a.log.Info("[Agent:Run] request failed: %v", err)
						time.Sleep(5 * time.Second)
					}

					if resp != nil {
						a.log.Info("[Agent:Run] received result from server: %v", resp)
					}
				}
			}
		}(i)
	}
}
