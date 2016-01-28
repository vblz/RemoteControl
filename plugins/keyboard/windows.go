// +build windows

package keyboard

import (
	// #include <wtypes.h>
	// #include <Winuser.h>
	"C"
	"syscall"
)
import "unsafe"

var (
	user32         = syscall.NewLazyDLL("user32.dll")
	keybdEventInfo = user32.NewProc("SendInput")
)

// for const bellow see https://msdn.microsoft.com/en-us/library/windows/desktop/ms646271(v=vs.85).aspx

// virtual key consts
const (
	vkSpace      = C.WORD(0x20)
	vkLeft       = 0x25
	vkUp         = 0x26
	vkRight      = 0x27
	vkDown       = 0x28
	vkVolumeDown = 0xAE
	vkVolumeUp   = 0xAF
)

// input type consts
const (
	inputTypeKeyboard = C.DWORD(1)
)

// input flags
const (
	keyEventFKeyUp = C.DWORD(2)
)

func press(key uint64) {
	switch key {
	case keySpace:
		keybdEvent(vkSpace)
	case keyLeft:
		keybdEvent(vkLeft)
	case keyUp:
		keybdEvent(vkUp)
	case keyRight:
		keybdEvent(vkRight)
	case keyDown:
		keybdEvent(vkDown)
	case keyVolDown:
		keybdEvent(vkVolumeDown)
	case keyVolUp:
		keybdEvent(vkVolumeUp)
	}
}

func keybdEvent(keyCode C.WORD) {
	downInput := C.KEYBDINPUT{wVk: keyCode}
	down := C.INPUT{_type: inputTypeKeyboard}
	var kiDown [32]byte
	copy(kiDown[:], C.GoBytes(unsafe.Pointer(&downInput), C.int(unsafe.Sizeof(downInput))))
	down.anon0 = kiDown

	upInput := C.KEYBDINPUT{wVk: keyCode, dwFlags: keyEventFKeyUp}
	up := C.INPUT{_type: inputTypeKeyboard}
	var kiUp [32]byte
	copy(kiUp[:], C.GoBytes(unsafe.Pointer(&upInput), C.int(unsafe.Sizeof(upInput))))
	up.anon0 = kiUp
	kis := []C.INPUT{
		down,
		up,
	}

	keybdEventInfo.Call(
		uintptr(2),
		uintptr(unsafe.Pointer(&kis[0])),
		uintptr(unsafe.Sizeof(down)),
	)
}
