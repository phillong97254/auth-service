package token

type IssuePasswordTokenInput struct {
	Username string
	Password string
	Scope    string // tuỳ chọn: xin bớt scope
}
