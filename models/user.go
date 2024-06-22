package models

type User struct {
	Id          string `json:"id"`
	OpenId      string `json:"open_id"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	StationId   string `json:"station_id"`
	StationName string `json:"station_name"`
	Phone       string `json:"phone"`
	PlateNumber string `json:"plate_number"`
	Avatar      string `json:"avatar"`
	Company     string `json:"company"`
	License     string `json:"license"`
	Status      int    `json:"status"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`

	PermissionList []string `json:"permission_list"`
}

type ListUserParam struct {
	UserIds     []string `json:"user_ids" query:"user_ids"`
	StationIds  []string `json:"station_ids" query:"station_ids"`
	Name        string   `json:"name" query:"name"`
	Types       []int    `json:"types" query:"types"`
	Phone       string   `json:"phone" query:"phone"`
	PlateNumber string   `json:"plate_number" query:"plate_number"`
	Page        int32    `json:"page"`
	PageSize    int32    `json:"page_size"`
}

type ListUserData struct {
	Users []*User `json:"users"`
	Total int64   `json:"total"`
}

type AddUserParam struct {
	User *User `json:"user"`
}

type UpdateUserParam struct {
	UserId      string  `json:"user_id"`
	Name        *string `json:"name"`
	StationId   *string `json:"station_id"`
	Phone       *string `json:"phone"`
	PlateNumber *string `json:"plate_number"`
	Company     *string `json:"company"`
	License     *string `json:"license"`
}

type DeleteUserParam struct {
	UserId string `json:"user_id"`
}
