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
