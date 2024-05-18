package models

type Station struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
	CreateTime string  `json:"create_time"`
	UpdateTime string  `json:"update_time"`
}

type ListStationParam struct {
	Ids  []int64 `json:"ids"`
	Name string  `json:"name"`
}

type ListStationData struct {
	Stations []*Station `json:"stations"`
}

type AddStationParam struct {
	Station *Station `json:"station"`
}

type DeleteStationParam struct {
	StationId int64 `json:"station_id"`
}

type UpdateStationParam struct {
	Station *Station `json:"station"`
}
