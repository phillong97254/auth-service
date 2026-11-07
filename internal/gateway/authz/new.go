package authz

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/phillong97254/auth-service/config"
	"github.com/phillong97254/auth-service/internal/signer"
	authzv1 "github.com/phillong97254/auth-service/pkg/gen/authz/v1"
)

type impl struct {
	conn    *grpc.ClientConn
	client  authzv1.AuthzServiceClient
	timeout time.Duration
}

// New tạo gRPC client tới authz-service. Mỗi lời gọi tự gắn token của CHÍNH auth-service
// (tự ký bằng signer, scope roles:read), không cần client secret.
// Production: dùng TLS/mTLS thay cho insecure.
func New(cfg *config.Config, s signer.Signer) (Gateway, error) {
	conn, err := grpc.NewClient(cfg.AuthzService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(serviceToken{signer: s, clientID: cfg.RestServer.Name}))
	if err != nil {
		return nil, err
	}
	return &impl{
		conn:    conn,
		client:  authzv1.NewAuthzServiceClient(conn),
		timeout: time.Duration(cfg.AuthzService.Timeout) * time.Second,
	}, nil
}

func (i *impl) Close() error { return i.conn.Close() }

// serviceToken: grpc.PerRPCCredentials gắn "authorization: Bearer <token của auth-service>".
type serviceToken struct {
	signer   signer.Signer
	clientID string
}

func (t serviceToken) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	out, err := t.signer.Sign(signer.SignInput{
		Subject: "service:" + t.clientID, Scope: "roles:read", ClientID: t.clientID,
	})
	if err != nil {
		return nil, err
	}
	return map[string]string{"authorization": "Bearer " + out.Token}, nil
}

func (serviceToken) RequireTransportSecurity() bool { return false }
