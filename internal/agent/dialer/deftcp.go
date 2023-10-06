package dialer

import (
	"context"
	"fmt"
	"net"

	"github.com/erupshis/yugpa_test/internal/helpers"
	"github.com/erupshis/yugpa_test/internal/logger"
	"github.com/erupshis/yugpa_test/internal/messages"
)

type DefaultTCP struct {
	serverAddr string
	log        logger.BaseLogger
}

func CreateDefaultTCP(serverAddr string, baseLogger logger.BaseLogger) BaseDialer {
	return &DefaultTCP{
		serverAddr: serverAddr,
		log:        baseLogger,
	}
}

func (d *DefaultTCP) MakeRequestToServer(ctx context.Context, request *messages.Request) (*messages.Response, error) {
	var dialer net.Dialer
	conn, err := dialer.DialContext(ctx, "tcp", d.serverAddr)
	if err != nil {
		return nil, fmt.Errorf("dial connection to server: %w", err)
	}
	defer helpers.ExecuteWithLogError(conn.Close, d.log)

	if err = d.writeRequestToConnection(*request, conn); err != nil {
		return nil, fmt.Errorf("request to server: %w", err)
	}

	var resp messages.Response
	if err = resp.Deserialize(conn); err != nil {
		return nil, fmt.Errorf("parse response from server: %w", err)
	}

	return &resp, nil
}

func (d *DefaultTCP) writeRequestToConnection(req messages.Request, conn net.Conn) error {
	requestBytes, err := req.Serialize()
	if err != nil {
		return fmt.Errorf("serialize request: %w", err)
	}

	_, err = conn.Write(requestBytes)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}

	return nil
}
