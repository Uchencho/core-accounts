package workflow

import (
	"time"

	"github.com/Uchencho/commons/uuid"
	"github.com/Uchencho/core-accounts/internal/db"
	"github.com/Uchencho/core-proto/generated/accounts"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateAccount(uuidGenerator uuid.GenV4Func,
	insertAccount db.InsertAccountFunc) CreateAccountFunc {
	return func(req accounts.CreateAccountRequest) error {
		a := accounts.Account{
			Id:        string(uuidGenerator()),
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Address:   req.Address,
			Age:       req.Age,
			CreatedAt: timestamppb.New(time.Now()),
		}

		if err := insertAccount(a); err != nil {
			return err
		}
		return nil
	}
}
