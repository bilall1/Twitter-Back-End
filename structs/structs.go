package structs

type UserFollower struct {
	UserId     int
	FollowerId int
}

type FollowingPeople struct {
	Id   int
	Page int
}
type FollowerPeople struct {
	Id   int
	Page int
}

type UserProfile struct {
	Id   int
	Link string
}

type Password struct {
	Id          int
	OldPassword string
	NewPassword string
}
type TweetOfUser struct {
	Email string
	Page  int
}

type TweetUser struct {
	TweetId int
	UserId  int
}
type TweetComment struct {
	TweetId int
	UserId  int
	Content string
}
type CommentLimit struct {
	TweetId int
	Limit   int
}

type TweetFollower struct {
	Id   int
	Page int
}

type CommentUser struct {
	Id           int
	TweetId      int
	UserId       int
	TweetComment string
	Email        string
	FirstName    string
	LastName     string
	Profile      string
}

type TweetData struct {
	Id        int
	Content   string
	UserId    int
	FirstName string
	LastName  string
	Email     string
	Profile   string
	Link      string
}
type ConversationData struct {
	UserId        int
	UserEmail     string
	UserFirstName string
	UserLastName  string
	UserProfile   string
	Id            int
	Participant1  int
	Participant2  int
	LastChat      string
	LastMessage   string
}

type Message struct {
	SenderId   int
	RecieverId int
	Content    string
	CreatedAt  string
}
