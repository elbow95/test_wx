package models

type Oil struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	StationId   int64  `json:"station_id"`
	StationName string `json:"station_name"`
	Price       int64  `json:"price"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

type ListOilParam struct {
	Ids        []int64 `json:"ids" query:"ids"`
	StationIds []int64 `json:"station_ids" query:"station_ids"`
	Name       string  `json:"name" query:"name"`
}

type ListOilData struct {
	Oils []*Oil `json:"oils"`
}

type AddOilParam struct {
	Oil *Oil `json:"oil"`
}

type UpdateOilParam struct {
	Oil *Oil `json:"oil"`
}

type DeleteOilParam struct {
	OilId int64 `json:"oil_id"`
}
