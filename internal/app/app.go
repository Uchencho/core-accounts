package app

import (
	"context"

	"github.com/Uchencho/commons/uuid"
	"github.com/Uchencho/core-accounts/internal/db"
	"github.com/Uchencho/core-accounts/internal/workflow"
	"github.com/Uchencho/core-proto/generated/accounts"
	"github.com/golang/protobuf/ptypes/empty"
)

type App struct {
	AppName string
	Option  Option
}

type Option struct {
	GenerateUUID  uuid.GenV4Func
	InsertAccount db.InsertAccountFunc
}

type OptionalArg func(oa *Option)

func NewApp(opts ...OptionalArg) App {
	o := Option{
		GenerateUUID: uuid.GenV4,
	}

	for _, opt := range opts {
		opt(&o)
	}

	return App{AppName: "core-accounts", Option: o}
}

func (a App) CreateAccount(c context.Context, acc *accounts.CreateAccountRequest) (*empty.Empty, error) {
	wf := workflow.CreateAccount(a.Option.GenerateUUID, a.Option.InsertAccount)
	if err := wf(*acc); err != nil {
		return nil, err
	}
	return nil, nil
}

func (a App) GetAccount(context.Context, *accounts.GetAccountRequest) (*accounts.Account, error) {
	return &accounts.Account{}, nil
}
