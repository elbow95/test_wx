package models

type Record struct {
	Id          string  `json:"id"`
	StationId   string  `json:"station_id"`
	StationName string  `json:"station_name"`
	OilId       string  `json:"oil_id"`
	OilName     string  `json:"oil_name"`
	StaffId     string  `json:"staff_id"`
	StaffName   string  `json:"staff_name"`
	StaffPhone  string  `json:"staff_phone"`
	Price       float64 `json:"price"`
	Liter       float64 `json:"liter"`
	Amount      float64 `json:"amount"`
	DriverId    string  `json:"driver_id"`
	DriverName  string  `json:"driver_name"`
	DriverPhone string  `json:"driver_phone"`
	CreateTime  string  `json:"create_time"`
	UpdateTime  string  `json:"update_time"`
}

type RangeInt64 struct {
	Left  int64
	Right int64
}

type ListRecordParam struct {
	Ids        []string    `json:"ids" query:"ids"`
	StationIds []string    `json:"station_ids" query:"station_ids"`
	OilIds     []string    `json:"oil_ids" query:"oil_ids"`
	StaffIds   []string    `json:"staff_ids" query:"staff_ids"`
	CreateTime *RangeInt64 `json:"create_time" query:"create_time"`
	Page       int32       `json:"page" query:"page"`
	PageSize   int32       `json:"page_size" query:"page_size"`
}

type ListRecordData struct {
	Records []*Record `json:"records"`
	Total   int64     `json:"total"`
}

type AddRecordParam struct {
	Record *Record `json:"record"`
}

type UpdateRecrodParam struct {
	RecordId    string   `json:"record_id"`
	StationId   *string  `json:"station_id"`
	OilId       *string  `json:"oil_id"`
	Price       *float64 `json:"price"`
	Liter       *float64 `json:"Liter"`
	Amount      *float64 `json:"amount"`
	DriverName  *string  `json:"driver_name"`
	DriverPhone *string  `json:"driver_phone"`
}

type DeleteRecordParam struct {
	RecordId string `json:"record_id"`
}
