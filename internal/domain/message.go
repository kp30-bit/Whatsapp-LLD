package domain

type Message struct {
	Id         int
	SenderId   int
	ReceiverId int
	Type       MessageType
	Content    string
}
