package workflow

import "github.com/Uchencho/core-proto/generated/accounts"

type CreateAccountFunc func(req accounts.CreateAccountRequest) error
