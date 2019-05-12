package rules

const (
	// 设备类型常量
	PMC350F = 1
	VJG     = 2
	VJL     = 3
	RCN350F = 4
	PMC350  = 5
	RCNVJ   = 6

	// 设备状态
	Inactive = 0 // 设备未激活
	Online   = 1 // 设备正常在线
	Offline  = 2 // 设备离线
	Warn     = 3 // 设备预警
	Alert    = 4 // 设备报警

	// Redis Key
	DeviceStatus   = `devices/status`
	ServiceVersion = `service/version`

	// RabbitMQ
	AlarmExchange = `alarms`
	ParamExchange = `params`

	// MongoDB
	DevicesCollection = `devices`

	// 服务类型

	ProductMaintenance = 1 // 产品维护
	AlarmProcessed     = 2 // 报警已处理
	Eliminated         = 3 // 排除电气隐患
	PendingEliminate   = 4 // 待排除电气隐患
	Other              = 5 // 其他
	Polling            = 6 // 巡检
)

var DeviceTypeLowerCaseFullName = map[int]string{
	PMC350F: "pmc350f",
	VJG:     "vjg",
	VJL:     "vjl",
	RCN350F: "rcn350f",
	PMC350:  "pmc350",
	RCNVJ:   "rcn-vj",
}

var DeviceTypeLowerCaseShortName = map[int]string{
	PMC350F: "f",
	RCN350F: "r",
	PMC350:  "a",
}

var DeviceStatusDESC = map[int]string{
	Inactive: "设备未激活",
	Online:   "设备正常",
	Offline:  "设备离线",
	Warn:     "设备预警",
	Alert:    "设备报警",
}

var TrueOrFalseDESC = map[bool]string{
	true:  "是",
	false: "否",
}

var MetricDESC = map[string]string{
	"DO1": "空开",
	"DO2": "报警器",
	"DI1": "烟雾报警",
	"DI2": "柜门状态",
	"Ia":  "A相电流",
	"Ib":  "B相电流",
	"Ic":  "C相电流",
	"IR":  "剩余电流",
	"T1":  "A相温度",
	"T2":  "B相温度",
	"T3":  "C相温度",
	"T4":  "箱体温度",
	"Ua":  "A相电压",
	"Ub":  "B相电压",
	"Uc":  "C相电压",
}
