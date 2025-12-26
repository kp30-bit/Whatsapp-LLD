package domain

type MessageType int

const (
	PersonalMessage MessageType = iota
	GroupMessage
)
