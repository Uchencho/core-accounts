package db

import "github.com/Uchencho/core-proto/generated/accounts"

type InsertAccountFunc func(accounts.Account) error
