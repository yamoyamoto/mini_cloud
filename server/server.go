package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/firecracker-microvm/firecracker-go-sdk"
	"github.com/firecracker-microvm/firecracker-go-sdk/client/models"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/create-instance", CreateInstance)
	e.Logger.Fatal(e.Start(":1323"))
}

func toPtr[T any](v T) *T {
	return &v
}

type CreateInstanceRequest struct {
}

type CreateInstanceResponse struct {
	IpAddress string `json:"ip_address"`
}

func CreateInstance(c echo.Context) error {
	now := time.Now().Unix()

	var vCpuCount int64 = 1
	var memSizeMib int64 = 128

	// TODO: 設定しっかり見直す
	cfg := firecracker.Config{
		SocketPath:      fmt.Sprintf("/tmp/firecracker_%d.socket", now),
		KernelImagePath: "../../vmlinux-5.10.186",
		LogPath:         fmt.Sprintf("%s.log", fmt.Sprintf("/tmp/firecracker_%d.socket", now)),
		KernelArgs:      "console=ttyS0 reboot=k panic=1 pci=off",
		MachineCfg: models.MachineConfiguration{
			MemSizeMib: &memSizeMib,
			VcpuCount:  &vCpuCount,
		},
		Drives: []models.Drive{
			{
				DriveID:      toPtr("rootfs"),
				IsRootDevice: toPtr(true),
				PathOnHost:   toPtr("../../ubuntu-22.04.ext4"),
				IsReadOnly:   toPtr(false),
			},
		},
		NetworkInterfaces: []firecracker.NetworkInterface{{
			// Use CNI to get dynamic IP
			CNIConfiguration: &firecracker.CNIConfiguration{
				NetworkName: "fcnet",
				IfName:      "veth0",
			},
		}},
	}

	m, err := firecracker.NewMachine(context.Background(), cfg)
	if err != nil {
		return fmt.Errorf("Failed to create new machine: %v", err)
	}

	if err := m.Start(context.Background()); err != nil {
		return fmt.Errorf("Failed to start machine: %v", err)
	}

	time.Sleep(5 * time.Second)

	if len(m.Cfg.NetworkInterfaces) == 0{
		return fmt.Errorf("Failed to get network interface")
	}
	
	log.Printf("Virtual machine has started. ip address: %s", m.Cfg.NetworkInterfaces[0].StaticConfiguration.IPConfiguration.IPAddr.IP)

	c.JSON(200, CreateInstanceResponse{
		IpAddress: m.Cfg.NetworkInterfaces[0].StaticConfiguration.IPConfiguration.IPAddr.IP.String(),
	})

	return nil
}
