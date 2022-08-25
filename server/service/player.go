package service

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/yoshietao/wolf/server/authz"
	"github.com/yoshietao/wolf/server/db/models"
)

func SelectSeat(ctx context.Context, db *sql.DB, seatId int) error {
	sessionUser := authz.GetLoggedInUser(ctx)
	if sessionUser == nil || !reflect.ValueOf(sessionUser).IsValid() {
		return fmt.Errorf("no logged in user found")
	}
	dbUser := sessionUser.(*models.User)
	dbPlayers := dbUser.R.UserIdPlayers

	if len(dbPlayers) > 1 {
		return fmt.Errorf("internal error, user is associated with multiple players")
	}
	// if player struct is not empty, the user already selected a seat
	if len(dbPlayers) == 1 {
		return fmt.Errorf("user already selected a seat")
	}

	// player info is null, create a new row based on the input for a game

	// TODO(Atao): configure number of player

	/* seatId, err := strconv.Atoi(seatString)
	if err != nil {
		fmt.Println("Error parsing seat_id into an integer")
	}*/

	if seatId < 1 || seatId > 12 {
		return fmt.Errorf("seat number invalid, must be between 1 to 12")
	}

	dbPlayer := &models.Player{
		UserId:      dbUser.ID,
		SeatId:      seatId,
		GameId:      1,
		CharacterId: 1,
		Status:      "",
		Result:      "",
	}

	err := dbPlayer.Insert(ctx, db, boil.Infer())
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("internal error, cannot store seat info. Error: %s", err.Error()))
	}

	fmt.Println("Successfully register a seat")
	return nil
}
