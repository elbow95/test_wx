package common

type UserType int

var (
	UserType_Admin  UserType = 1
	UserType_Staff  UserType = 10
	UserType_Driver UserType = 20
)

type UserStatus int

var (
	UserStatus_Valid      UserStatus = 1
	UserStatus_WaitVerify UserStatus = 2
)
