package app

import (
	"context"
)

const (
	DEFAULT_KERNEL    = "./vmlinux-5.10.186"
	DEFAULT_BOOT_ARGS = "console=ttyS0 reboot=k panic=1 pci=off"
	DEFAULT_ROOTFS    = "./ubuntu-22.04.ext4"
)

func runCreateInstanceCmd(ctx context.Context, host string, instanceIp string) error {
	// TODO: implement

	return nil
}
