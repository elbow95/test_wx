package models

type User struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	StationiId  int64  `json:"station_id"`
	Phone       string `json:"phone"`
	PlateNumber string `json:"plate_number"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

type ListUserParam struct {
	Ids         []int64 `json:"ids"`
	StationIds  []int64 `json:"station_ids"`
	Name        string  `json:"name"`
	Types       []int   `json:"types"`
	Phone       string  `json:"phone"`
	PlateNumber string  `json:"plate_number"`
}

type ListUserData struct {
	Users []*User `json:"users"`
}

type AddUserParam struct {
	User *User `json:"user"`
}

type UpdateUserParam struct {
	User *User `json:"user"`
}

type DeleteUserParam struct {
	UserId int64 `json:"user_id"`
}
