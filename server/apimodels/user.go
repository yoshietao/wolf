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

type Character struct {
	Type CharacterType
	Team TeamType
}

type UserInput struct {
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
}
