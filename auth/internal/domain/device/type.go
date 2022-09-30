package device

import (
	"auth/internal/domain/device/agent"
	"auth/internal/domain/device/deviceID"
	"auth/internal/domain/device/ip"
	"auth/internal/domain/device/refreshToken"
	"time"

	"github.com/google/uuid"
)

type Device struct {
	id         uuid.UUID
	createdAt  time.Time
	modifiedAt time.Time

	userID       uuid.UUID
	deviceID     deviceID.DeviceID
	ip           ip.Ip
	agent        agent.Agent
	dtype        DeviceType
	refreshToken refreshToken.RefreshToken
	refreshExp   time.Time
	lastSeen     time.Time
}

func NewWithID(
	id uuid.UUID,
	createdAt time.Time,
	modifiedAt time.Time,

	userID uuid.UUID,
	deviceID deviceID.DeviceID,
	ip ip.Ip,
	agent agent.Agent,
	dtype DeviceType,
	refreshToken refreshToken.RefreshToken,
	refreshExp time.Time,
	lastSeen time.Time,
) (*Device, error) {
	return &Device{
		id:           id,
		createdAt:    createdAt,
		modifiedAt:   modifiedAt,
		userID:       userID,
		deviceID:     deviceID,
		ip:           ip,
		agent:        agent,
		dtype:        dtype,
		refreshToken: refreshToken,
		refreshExp:   refreshExp,
		lastSeen:     lastSeen,
	}, nil
}

func New(
	userID uuid.UUID,
	deviceID deviceID.DeviceID,
	ip ip.Ip,
	agent agent.Agent,
	dtype DeviceType,
	refreshToken refreshToken.RefreshToken,
	refreshExp time.Time,
	lastSeen time.Time,
) (*Device, error) {
	timeNow := time.Now()
	return &Device{
		id:           uuid.New(),
		createdAt:    timeNow,
		modifiedAt:   timeNow,
		userID:       userID,
		deviceID:     deviceID,
		ip:           ip,
		agent:        agent,
		dtype:        dtype,
		refreshToken: refreshToken,
		refreshExp:   refreshExp,
		lastSeen:     lastSeen,
	}, nil
}

func (d Device) ID() uuid.UUID {
	return d.id
}

func (d Device) CreatedAt() time.Time {
	return d.createdAt
}

func (d Device) ModifiedAt() time.Time {
	return d.modifiedAt
}

func (d Device) UserID() uuid.UUID {
	return d.userID
}

func (d Device) DeviceID() deviceID.DeviceID {
	return d.deviceID
}

func (d Device) Ip() ip.Ip {
	return d.ip
}

func (d Device) Agent() agent.Agent {
	return d.agent
}

func (d Device) Type() DeviceType {
	return d.dtype
}

func (d Device) RefreshToken() refreshToken.RefreshToken {
	return d.refreshToken
}

func (d Device) RefreshExpiration() time.Time {
	return d.refreshExp
}

func (d Device) LastSeen() time.Time {
	return d.lastSeen
}
