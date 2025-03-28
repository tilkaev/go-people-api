package model

type Friendship struct {
	ID       uint
	UserID   uint
	FriendID uint
}

type Person struct {
	ID          uint
	FirstName   string
	LastName    string
	Gender      string
	Nationality string
	Age         int
}
