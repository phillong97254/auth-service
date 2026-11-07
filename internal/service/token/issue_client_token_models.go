package token

type IssueClientTokenInput struct {
	ClientID     string
	ClientSecret string
	Scope        string // tuỳ chọn: xin bớt scope
}
