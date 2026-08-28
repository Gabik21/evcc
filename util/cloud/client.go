package cloud

import (
	_ "embed"
	"time"

	"github.com/evcc-io/evcc/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	hostport = util.Getenv("GRPC_URI", "sponsor.evcc.io:8080")

	conn *grpc.ClientConn
)

func Connection() (*grpc.ClientConn, error) {
	if conn != nil {
		return conn, nil
	}

	// close idle connection shortly after startup auth instead of churning against the server's idle close
	conn, err := grpc.NewClient(hostport,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithIdleTimeout(5*time.Second),
	)

	return conn, err
}
