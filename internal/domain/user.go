package domain

type User struct {
	Id          int
	Name        string
	ReceivedMsg map[int]*Message
}
