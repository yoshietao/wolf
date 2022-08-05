package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/yoshietao/wolf/server/apimodels"
	"github.com/yoshietao/wolf/server/db/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx context.Context, db *sql.DB, UserInput apimodels.UserInput) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(UserInput.PassWord), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	uuid := uuid.New()

	dbUser := models.User{
		UUID:         uuid.String(),
		UserName:     UserInput.UserName,
		PasswordHash: string(hash),
	}

	err = dbUser.Insert(ctx, db, boil.Infer())
	if err != nil {
		fmt.Println("cannot insert user.")
	}

	return uuid.String(), nil
}
