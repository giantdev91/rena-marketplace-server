package token

import (
	tokenModel "rena-server-v2/internal/token/model"
)

type TokenResponse struct {
	Token Token `json:"token"`
}

type TokensResponse struct {
	Tokens      []Token `json:"tokens"`
	TokensCount int64   `json:"tokensCount"`
}

type Token struct {
	ID    uint   `json:"id"`
	Token string `json:"token"`
}

// NewTokensResponse converts token models and total count to TokensResponse
func NewTokensResponse(tokens []*tokenModel.Token, total int64) *TokensResponse {
	var a []Token
	for _, token := range tokens {
		a = append(a, NewTokenResponse(token).Token)
	}

	return &TokensResponse{
		Tokens:      a,
		TokensCount: total,
	}
}

// NewTokenResponse converts token model to TokenResponse
func NewTokenResponse(a *tokenModel.Token) *TokenResponse {
	return &TokenResponse{
		Token: Token{
			ID:    a.ID,
			Token: a.Token,
		},
	}
}
