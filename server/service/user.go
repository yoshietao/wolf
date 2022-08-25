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

func CreateUser(ctx context.Context, db *sql.DB, UserInput apimodels.UserInput) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(UserInput.PassWord), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	dbUser, err := models.Users(models.UserWhere.UserName.EQ(UserInput.UserName)).One(ctx, db)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("cannot retrieve user info: %s", err.Error())
	} else if dbUser != nil {
		return fmt.Errorf("user name %s already exists in the databse, try login", UserInput.UserName)
	}

	uuid := uuid.New()

	dbUser = &models.User{
		UUID:         uuid.String(),
		UserName:     UserInput.UserName,
		PasswordHash: string(hash),
	}

	err = dbUser.Insert(ctx, db, boil.Infer())
	if err != nil {
		fmt.Println("cannot insert user.")
	}

	return nil
}

// TODO: consider return only error for either username/password error
// As returning separate errors is an anti-pattern
func LoginUser(ctx context.Context, db *sql.DB, UserInput apimodels.UserInput) (string, error) {
	dbUser, err := models.Users(models.UserWhere.UserName.EQ(UserInput.UserName)).One(ctx, db)
	if err != nil {
		return "", fmt.Errorf("cannot retrieve user info: %s", err.Error())
	} else if dbUser == nil {
		return "", fmt.Errorf("user %s does not exist, try creat a new account", err.Error())
	}

	if err = (bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(UserInput.PassWord))); err != nil {
		return "", fmt.Errorf("password is incorrect")
	}

	fmt.Println(fmt.Sprintf("user %s logged in", dbUser.UUID))

	return dbUser.UUID, nil
}
