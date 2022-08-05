package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/yoshietao/wolf/server/apimodels"
	"github.com/yoshietao/wolf/server/db/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx context.Context, db *sql.DB, UserInput apimodels.UserInput) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(UserInput.PassWord), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	dbUser := models.User{
		UserName:     null.StringFrom(UserInput.UserName),
		PasswordHash: null.StringFrom(string(hash)),
	}

	err = dbUser.Insert(ctx, db, boil.Infer())
	if err != nil {
		fmt.Println("cannot insert user.")
	}

	return nil
}
