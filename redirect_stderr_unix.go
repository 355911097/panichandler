// copy of https://github.com/ncw/rclone/blob/006227baed4fb53413b0598170d6c5cabf4e350e/fs/redirect_stderr_unix.go
// Log the panic under unix to the log file

// +build !windows,!solaris,!plan9

package panichandler

import (
	"log"
	"os"

	"golang.org/x/sys/unix"
)

// redirectStderr to the file passed in
func redirectStderr(f *os.File) {
	err := unix.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		log.Fatalf("Failed to redirect stderr to file: %v", err)
	}
}
