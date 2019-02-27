package rules

const (
	// 设备状态
	Inactive = 0 // 设备未激活
	Online   = 1 // 设备正常在线
	Offline  = 2 // 设备离线
	Warn     = 3 // 设备预警
	Alert    = 4 // 设备报警

	// Redis Key
	DeviceStatus = `devices/status`
)
