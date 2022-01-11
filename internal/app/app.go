package app

import (
	"context"
	"log"

	"github.com/Uchencho/commons/uuid"
	"github.com/Uchencho/core-accounts/internal/db"
	"github.com/Uchencho/core-accounts/internal/workflow"
	"github.com/Uchencho/core-proto/generated/accounts"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
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
	log.Printf("received create account: %+v", acc)
	wf := workflow.CreateAccount(a.Option.GenerateUUID, a.Option.InsertAccount)
	if err := wf(*acc); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (a App) GetAccount(context.Context, *accounts.GetAccountRequest) (*accounts.Account, error) {
	return &accounts.Account{}, nil
}
