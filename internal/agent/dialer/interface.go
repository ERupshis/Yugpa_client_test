package dialer

import (
	"context"

	"github.com/erupshis/yugpa_test/internal/messages"
)

//go:generate mockgen -destination=../../../mocks/mock_BaseDialer.go -package=mocks github.com/erupshis/yugpa_test/internal/agent/dialer BaseDialer
type BaseDialer interface {
	MakeRequestToServer(ctx context.Context, request *messages.Request) (*messages.Response, error)
}
