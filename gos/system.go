package gos

import (
	"os"
	"runtime"
	"strings"
)

// IsWin determine whether the system is windows
func IsWin() bool {
	return runtime.GOOS == "windows"
}

// IsMac determines whether the system is darwin
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// IsLinux determines whether the system is linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsSupportColor check current console whether support color.
// Supported: linux, mac, or windows's ConEmu, Cmder, putty, git-bash.exe
// Not support: windows cmd.exe, powerShell.exe
func IsSupportColor() bool {
	// Support color: "TERM=xterm" "TERM=xterm-vt220" "TERM=xterm-256color" "TERM=screen-256color"
	// Don't support color: "TERM=cygwin"
	envTerm := os.Getenv("TERM")
	if strings.Contains(envTerm, "xterm") || strings.Contains(envTerm, "screen") {
		return true
	}

	// like on ConEmu software, e.g "ConEmuANSI=ON"
	if os.Getenv("ConEmuANSI") == "ON" {
		return true
	}

	// like on ConEmu software, e.g "ANSICON=189x2000 (189x43)"
	if os.Getenv("ANSICON") != "" {
		return true
	}

	return false
}

// IsSupport256Color check current console whether support 256 color.
func IsSupport256Color() bool {
	// "TERM=xterm-256color" "TERM=screen-256color"
	return strings.Contains(os.Getenv("TERM"), "256color")
}

// IsSupportTrueColor check current console whether support true color
func IsSupportTrueColor() bool {
	// "COLORTERM=truecolor"
	return strings.Contains(os.Getenv("COLORTERM"), "truecolor")
}
