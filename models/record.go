package models

type Record struct {
	Id          int64  `json:"id"`
	StationId   int64  `json:"station_id"`
	OilId       int64  `json:"oil_id"`
	StaffId     int64  `json:"staff_id"`
	Price       int64  `json:"price"`
	Liter       int64  `json:"liter"`
	Amount      int64  `json:"amount"`
	DriverId    int64  `json:"driver_id"`
	DriverName  string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

type RangeInt64 struct {
	Left  int64
	Right int64
}

type ListRecordParam struct {
	Ids        []int64     `json:"ids"`
	StationIds []int64     `json:"station_ids"`
	OilIds     []int64     `json:"oil_ids"`
	StaffIds   []int64     `json:"staff_ids"`
	CreateTime *RangeInt64 `json:"create_time"`
}

type ListRecordData struct {
	Records []*Record `json:"records"`
}

type AddRecordParam struct {
	Record *Record `json:"record"`
}

type DeleteRecordParam struct {
	RecordId int64 `json:"record_id"`
}
