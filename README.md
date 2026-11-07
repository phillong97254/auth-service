# auth-service  (:9000)

OAuth2 Authorization Server tối giản (thay cho Keycloak): `POST /oauth/token` (password, client_credentials) và `GET /.well-known/jwks.json`.

Chỉ kiểm tra **ai là ai** (tài khoản + mật khẩu, client secret). **Role/scope** của user lấy từ
**authz-service** (`GetUserRoles`, gRPC) ngay lúc phát token; auth-service tự ký token m2m (scope `roles:read`)
để gọi, nên không cần client secret. authz-service không trả lời được thì login trả `503 temporarily_unavailable`
(không phát token thiếu quyền). Địa chỉ: `authz_service.addr` / env `AUTHZ_SERVICE_ADDR`.

Các service khác lấy JWKS từ đây để verify token.

## Chạy

Folder này là 1 project Go độc lập (repo git, module riêng), mở folder lên rồi:

```bash
make env              # lần đầu: tạo .env từ .env.example
go run ./cmd/rest    # hoặc: make run
make test             # go test -race ./...
```

Cấu hình ở `config/config.yml`, ghi đè bằng biến môi trường theo quy tắc `.` → `_`
(ví dụ `REST_SERVER_ADDR=:9999`, `OAUTH_JWKS_URL=...`).

## Secret và cấu hình

- `config/config.yml` **được commit**: chỉ chứa giá trị mặc định, không chứa secret.
- Secret đặt trong `.env` (đã có trong `.gitignore` và `.dockerignore`, không commit, không vào image).
  `make env` tạo `.env` từ `.env.example`, sau đó điền giá trị. `make run` tự nạp `.env`.
- Thiếu secret thì service dừng ngay lúc khởi động và báo tên biến cần đặt.
- Trên server/CI: truyền secret bằng biến môi trường (GitHub Actions secrets, Docker/K8s secret...), không dùng `.env` commit lên.

## Thư viện (make mod)

Code dùng chung (`authkit`, `apperr`) nằm ở repo **go-lib** (`github.com/phillong97254/go-lib`), giống `go-library`
ở các dự án trước. Thư viện được vendor vào `vendor/` nên build/Docker không cần mạng hay repo khác.

```bash
make mod                          # go get -u thư viện + go mod tidy + go mod vendor
make mod GO_LIB_VERSION=v0.2.0    # chốt go-lib ở 1 version (khi dùng remote)
make mod-local                    # dùng ../go-lib (replace) khi đang sửa go-lib cùng lúc
make mod-remote                   # bỏ replace, lấy go-lib theo version từ GitHub (mặc định hiện tại)
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
    internal/signer                      # giữ private key, ký JWT RS256, JWKS
    internal/gateway/authz               # gRPC client tới authz-service
    internal/service/token               # phát token (password: hỏi role ở authz-service; client_credentials)
    internal/handler/rest/v1/public/oauth
```

## gRPC contract

`api/proto/authz/v1/authz.proto` → `pkg/gen/authz/v1`. Đây là bản copy hợp đồng của authz-service: sửa proto ở
authz-service trước rồi copy file `.proto` sang đây, sau đó `make gen-proto`.
