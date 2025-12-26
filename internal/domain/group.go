package domain

type Group struct {
	Id          int
	Name        string
	Members     []*User
	ReceivedMsg map[int]*Message
}
