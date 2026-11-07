# auth-service  (:9000)

OAuth2 Authorization Server tối giản (thay cho Keycloak): `POST /oauth/token` (password, client_credentials) và `GET /.well-known/jwks.json`.

Không phụ thuộc service nào. Phải chạy **trước** vì các service khác lấy JWKS từ đây.

## Chạy

Folder này là 1 project Go độc lập (module riêng, không cần `go.work`), mở folder lên rồi:

```bash
go run ./cmd/rest    # hoặc: make run
make test             # go test -race ./...
```

Cấu hình ở `config/config.yml`, ghi đè bằng biến môi trường theo quy tắc `.` → `_`
(ví dụ `REST_SERVER_ADDR=:9999`, `OAUTH_JWKS_URL=...`).

## Docker

```bash
make build            # docker build -t go-admin-wallet/auth-service:dev-latest .
make up               # chạy riêng service này (gọi service khác qua host.docker.internal)
make ci-docker-build ci-docker-push REGISTRY=<registry>/<group> TAG=v1.0.0
```

## Cấu trúc

```
cmd/rest/main.go, routes/   # bootstrap + đăng ký route
config/                      # config.go (viper) + config.yml
pkg/authkit                  # verify JWT (JWKS), middleware REST, interceptor gRPC
internal/
    domain/
    internal/repository/{user,client}   # dữ liệu demo (in-memory)
    internal/service/token               # ký JWT RS256, JWKS, thu hẹp scope
    internal/handler/rest/v1/public/oauth
```
