# auth-service  (:9000)

OAuth2 Authorization Server tối giản (thay cho Keycloak): `POST /oauth/token` (password, client_credentials) và `GET /.well-known/jwks.json`.

Không phụ thuộc service nào. Phải chạy **trước** vì các service khác lấy JWKS từ đây.

## Chạy

Folder này là 1 project Go độc lập (repo git, module riêng), mở folder lên rồi:

```bash
go run ./cmd/rest    # hoặc: make run
make test             # go test -race ./...
```

Cấu hình ở `config/config.yml`, ghi đè bằng biến môi trường theo quy tắc `.` → `_`
(ví dụ `REST_SERVER_ADDR=:9999`, `OAUTH_JWKS_URL=...`).

## Thư viện (make mod)

Code dùng chung (`authkit`, `apperr`) nằm ở repo **go-lib** (`github.com/phillong97254/go-lib`), giống `go-library`
ở các dự án trước. Thư viện được vendor vào `vendor/` nên build/Docker không cần mạng hay repo khác.

```bash
make mod                          # go get -u thư viện + go mod tidy + go mod vendor
make mod GO_LIB_VERSION=v0.2.0    # chốt go-lib ở 1 version (khi dùng remote)
make mod-local                    # dùng ../go-lib (replace) – mặc định hiện tại vì go-lib chưa push
make mod-remote                   # bỏ replace, lấy go-lib theo version từ remote (cần make netrc nếu repo private)
```

## Docker

```bash
make build            # docker build -t ghcr.io/phillong97254/auth-service:dev-latest .
make up               # chạy riêng service này (gọi service khác qua host.docker.internal)
make ci-docker-build ci-docker-push TAG=v1.0.0   # push lên ghcr.io/phillong97254 (GitHub Container Registry)
```

## Cấu trúc

```
cmd/rest/main.go, routes/   # bootstrap + đăng ký route
config/                      # config.go (viper) + config.yml
vendor/                      # thư viện (kể cả go-lib), tạo bằng make mod
internal/
    domain/
    internal/repository/{user,client}   # dữ liệu demo (in-memory)
    internal/service/token               # ký JWT RS256, JWKS, thu hẹp scope
    internal/handler/rest/v1/public/oauth
```
