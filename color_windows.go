package color

import (
	"os"
	"runtime"
	"syscall"
)

func init() {
	if runtime.GOOS == "windows" {
		// Try to make ANSI work
		handle := syscall.Handle(os.Stdout.Fd())
		kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
		setConsoleModeProc := kernel32DLL.NewProc("SetConsoleMode")
		// If it fails, fallback to no colors
		if _, _, err := setConsoleModeProc.Call(uintptr(handle), 0x0001|0x0002|0x0004); err != nil && err.Error() != "The operation completed successfully." {
			Reset = ""
			Bold = ""
			Underline = ""
			Black = ""
			Red = ""
			Green = ""
			Yellow = ""
			Blue = ""
			Purple = ""
			Cyan = ""
			Gray = ""
			White = ""
			BlackBackground = ""
			RedBackground = ""
			GreenBackground = ""
			YellowBackground = ""
			BlueBackground = ""
			PurpleBackground = ""
			CyanBackground = ""
			GrayBackground = ""
			WhiteBackground = ""
		}
	}
}
