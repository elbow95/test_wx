package common

var (
	PermissionJingangSaoMa  = "jingang_saoma" // 扫码权限
	PermissionJingangRecord = "record_list"   // 加油记录列表
	PermissionJingangOil    = "oil_list"      // 油品油价列表
	PermissionJingangUser   = "user_list"     // 用户列表
	PermissionJingangDriver = "driver_list"   // 司机列表

	PermissionStationAdd    = "station_add"    // 油站添加
	PermissionStationDelete = "station_delete" // 油站删除
	PermissionStationUpdate = "station_update" // 油站更新

	PermissionOilAdd    = "oil_add"    // 油品增加
	PermissionOilUpdate = "oil_update" // 油品更新
	PermissionOilDelete = "oil_delete" // 油品删除
)

var (
	UserPermissionMap = map[UserType][]string{
		UserType_Admin: {
			PermissionJingangSaoMa,
			PermissionJingangRecord,
			PermissionJingangOil,
			PermissionJingangUser,
			PermissionJingangDriver,
			PermissionStationAdd,
			PermissionStationDelete,
			PermissionStationUpdate,
			PermissionOilAdd,
			PermissionOilUpdate,
			PermissionOilDelete,
		},
		UserType_Staff: {
			PermissionJingangSaoMa,
			PermissionJingangRecord,
			PermissionJingangOil,
			PermissionOilAdd,
			PermissionOilUpdate,
			PermissionOilDelete,
		},
		UserType_Driver: {
			PermissionJingangSaoMa,
			PermissionJingangOil,
		},
	}
)
