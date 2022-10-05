package device

import (
	"auth/internal/domain/device"
	"auth/internal/domain/device/agent"
	"auth/internal/domain/device/deviceID"
	"auth/internal/domain/device/ip"
	"auth/internal/domain/device/refreshToken"
	"auth/internal/repository/postgres/device/dao"
)

func (r Repository) toDomainDevice(dao *dao.Device) (result *device.Device, err error) {
	agent, err := agent.NewAgent(dao.Agent)
	if err != nil {
		return
	}

	deviceID, err := deviceID.NewDeviceID(dao.DeviceID)
	if err != nil {
		return
	}

	ip, err := ip.NewIp(dao.Ip)
	if err != nil {
		return
	}

	refreshToken, err := refreshToken.NewFromString(dao.RefreshToken)
	if err != nil {
		return
	}

	result, err = device.NewWithID(
		dao.ID,
		dao.CreatedAt,
		dao.ModifiedAt,
		dao.UserID,
		*deviceID,
		*ip,
		*agent,
		device.DeviceType(dao.Dtype),
		*refreshToken,
		dao.RefreshExp,
		dao.LastSeen,
	)

	return
}
