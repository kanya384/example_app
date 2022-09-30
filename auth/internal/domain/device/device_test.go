package device

import (
	"auth/internal/domain/device/agent"
	"auth/internal/domain/device/deviceID"
	"auth/internal/domain/device/ip"
	"auth/internal/domain/device/refreshToken"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewWithID(t *testing.T) {
	req := require.New(t)
	id := uuid.New()
	timeNow := time.Now()
	deviceID, _ := deviceID.NewDeviceID("00000000-000000000000000")
	userID := uuid.New()
	ip, _ := ip.NewIp("8.8.8.8")
	refreshToken, _ := refreshToken.New()
	agent, _ := agent.NewAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 Edge/15.15063")
	dtype := IOS
	t.Run("create company with id success", func(t *testing.T) {
		device, err := NewWithID(id, timeNow, timeNow, userID, *deviceID, *ip, *agent, dtype, *refreshToken, timeNow, timeNow)
		req.Equal(err, nil)
		req.Equal(device.ID(), id)
		req.Equal(device.CreatedAt(), timeNow)
		req.Equal(device.ModifiedAt(), timeNow)
		req.Equal(device.UserID(), userID)
		req.Equal(device.DeviceID(), *deviceID)
		req.Equal(device.Ip(), *ip)
		req.Equal(device.Agent(), *agent)
		req.Equal(device.Type(), dtype)
		req.Equal(device.RefreshToken(), *refreshToken)
		req.Equal(device.RefreshExpiration(), timeNow)
		req.Equal(device.LastSeen(), timeNow)
	})
}

func TestNew(t *testing.T) {
	req := require.New(t)
	timeNow := time.Now()
	deviceID, _ := deviceID.NewDeviceID("00000000-000000000000000")
	userID := uuid.New()
	ip, _ := ip.NewIp("8.8.8.8")
	refreshToken, _ := refreshToken.New()
	agent, _ := agent.NewAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 Edge/15.15063")
	dtype := IOS
	t.Run("create company with id success", func(t *testing.T) {
		device, err := New(userID, *deviceID, *ip, *agent, dtype, *refreshToken, timeNow, timeNow)
		req.Equal(err, nil)
		req.NotEmpty(device.ID())
		req.NotEmpty(device.CreatedAt())
		req.NotEmpty(device.ModifiedAt())
		req.Equal(device.UserID(), userID)
		req.Equal(device.DeviceID(), *deviceID)
		req.Equal(device.Ip(), *ip)
		req.Equal(device.Agent(), *agent)
		req.Equal(device.Type(), dtype)
		req.Equal(device.RefreshToken(), *refreshToken)
		req.Equal(device.RefreshExpiration(), timeNow)
		req.Equal(device.LastSeen(), timeNow)
	})
}
