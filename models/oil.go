package models

type Oil struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	StationId   string  `json:"station_id"`
	StationName string  `json:"station_name"`
	Price       float64 `json:"price"`
	CreateTime  string  `json:"create_time"`
	UpdateTime  string  `json:"update_time"`
}

type ListOilParam struct {
	Ids        []string `json:"ids" query:"ids"`
	StationIds []string `json:"station_ids" query:"station_ids"`
	Name       string   `json:"name" query:"name"`
}

type ListOilData struct {
	Oils []*Oil `json:"oils"`
}

type AddOilParam struct {
	Oil *Oil `json:"oil"`
}

type UpdateOilParam struct {
	OilId string   `json:"oil_id"`
	Price *float64 `json:"price"`
}

type DeleteOilParam struct {
	OilId string `json:"oil_id"`
}

type BatchUpdateOilPriceParam struct {
	Oils []*Oil `json:"oils"`
}
