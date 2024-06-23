package models

type Station struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Mobile     string  `json:"mobile"`
	Address    string  `json:"address"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
	CreateTime string  `json:"create_time"`
	UpdateTime string  `json:"update_time"`
}

type ListStationParam struct {
	Ids  []string `json:"ids" query:"ids"`
	Name string   `json:"name" query:"name"`
}

type ListStationData struct {
	Stations []*Station `json:"stations"`
	Total    int64      `json:"total"`
}

type AddStationParam struct {
	Station *Station `json:"station"`
}

type DeleteStationParam struct {
	StationId string `json:"station_id"`
}

type UpdateStationParam struct {
	StationId string   `json:"station_id"`
	Mobile    *string  `json:"mobile"`
	Name      *string  `json:"name"`
	Address   *string  `json:"address"`
	Longitude *float64 `json:"longitude"`
	Latitude  *float64 `json:"latitude"`
}
