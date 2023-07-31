package models

type User struct {
	Id         int
	Email      string
	Password   string
	ThirdParty bool
	D_o_b      string
	FirstName  string
	LastName   string
}

type Tweet struct {
	Id      int
	Content string
	UserId  int
}

type UserFollowers struct {
	Id        int
	UserId    int
	FolloweId int
}