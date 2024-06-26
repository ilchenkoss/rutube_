package port

//go:generate mockgen -source=./telegram.go -destination=mock/telegram.go -package=mock

type Telegram interface {
	GetInviteLink(chatID int64, birthdayUsernames string) (string, error)
	KickUser(chatID int64, userID int64) error
	UnBanUser(chatID int64, userID int64) error
	SendMessage(chatID int64, text string)
}
