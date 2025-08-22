//go:build windows

package main

import (
	"context"
	"os/exec"
)

func ShellContext(ctx context.Context, cmd string) error {
	if err := exec.CommandContext(ctx, "C:\\Windows\\System32\\cmd.exe", "/c", cmd).Run(); err != nil {
		return err
	}
	return nil
}
