//go:build linux

package main

import (
	"context"
	"fmt"
	"os/exec"
)

func ShellContext(ctx context.Context, cmd string) error {
	fmt.Printf("test:")
	if err := exec.CommandContext(ctx, "/bin/bash", "-c", fmt.Sprintf("wine %s", cmd)).Run(); err != nil {
		return err
	}
	return nil
}
