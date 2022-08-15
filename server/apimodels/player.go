package apimodels

const (
	PlayerStatusTypeAlive PlayerStatusType = iota
	PlayerStatusTypeDead
)

type PlayerInput struct {
	// SeatId is the seat number, it's ranging from 1 to 12
	SeatId    int `json:"seat_id"`
	Character Character
	// If a player is a sheriff, this can be changed as the game goes
	Sheriff bool
	Status  PlayerStatusType

	// Additional attributed WIP
	// TODO: if we have guard
	Protected bool

	// result
	Win bool
}
