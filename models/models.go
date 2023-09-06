package models

type User struct {
	Id         int
	Email      string
	Password   string
	ThirdParty bool
	D_o_b      string
	FirstName  string
	LastName   string
	Profile    string
}

type Tweet struct {
	Id      int
	Content string
	UserId  int
	Link    string
}

type UserFollowers struct {
	Id        int
	UserId    int
	FolloweId int
}

type TweetLikes struct {
	Id      int
	TweetId int
	UserId  int
}

type Message struct {
	Id          int
	SenderId    int
	RecieverId  int
	MessageType string
	CreatedAt   string
	Status      string
	Content     string
}

type Conversation struct {
	Id           int
	Participant1 int
	Participant2 int
	LastChat     string
	LastMessage  string
}
