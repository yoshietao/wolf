package apimodels

type TeamType int
type CharacterType int
type PlayerStatusType int

// TeamType is used to determine the game result.
const (
	TeamTypeVillager TeamType = iota
	TeamTypeWerewolf
)

const (
	// There will be 4 villagers
	CharacterTypeVillager1 CharacterType = iota
	CharacterTypeVillager2
	CharacterTypeVillager3
	CharacterTypeVillager4

	// There will be 4 gods
	CharacterTypeSeer
	CharacterTypeWitch
	CharacterTypeHunter
	CharacterTypeIdiot

	// There will be 4 werewolfs
	CharacterTypeWerewolf1
	CharacterTypeWerewolf2
	CharacterTypeWerewolf3
	CharacterTypeWerewolf4
)

const (
	PlayerStatusTypeAlive PlayerStatusType = iota
	PlayerStatusTypeDead
)

type Player struct {
	Id int
	// Player name, could be used to calculate player status
	Name string `json:"name"`
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

type Character struct {
	Type CharacterType
	Team TeamType
}

type UserInput struct {
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
}
