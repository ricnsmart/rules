package rules

const (
	// 设备类型常量
	PMC350F = 1
	VJG     = 2
	VJL     = 3
	RCN350F = 4
	PMC350  = 5

	// 设备状态
	Inactive     = 0 // 设备未激活
	Online       = 1 // 设备正常在线
	Offline      = 2 // 设备离线
	Warn         = 3 // 设备预警
	Alert        = 4 // 设备报警
	Unregistered = 5 // 设备未注册

	// Redis Key
	DeviceStatus   = `devices/status`
	ServiceVersion = `service/version`

	// RabbitMQ
	AlarmExchange = `alarm`
	ParamExchange = `param`

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
}

var DeviceTypeLowerCaseShortName = map[int]string{
	PMC350F: "f",
	RCN350F: "r",
	PMC350:  "a",
}
