-- Tài khoản demo (mật khẩu: alice123 / bob123 / carol123, lưu dạng bcrypt).
-- Role/scope KHÔNG ở đây: authz-service quản lý.
INSERT INTO users (username, password_hash) VALUES
    ('alice', '$2a$10$BnFCSVj4cHiArEJAcbHs.uUItVWBcWO8STgkAhu1oLp5BjmBe5eAC'),
    ('bob',   '$2a$10$MCiy.YkELGXuCvdGyOH7DOIoAryf3cEwvlni/jG/l0ELklplRVs1O'),
    ('carol', '$2a$10$U4y9NeN0HGqbsEzDqC7.z.gRz89AnIR95/0f171yHYdwLRqw4jvDK')
ON CONFLICT (username) DO NOTHING;
