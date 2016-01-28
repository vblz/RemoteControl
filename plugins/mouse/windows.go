// +build windows

package mouse

import (
	"syscall"
	"unsafe"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	setCursorPosInfo = user32.NewProc("SetCursorPos")
	getCursorPosInfo = user32.NewProc("GetCursorPos")
	mouseEventInfo   = user32.NewProc("mouse_event")
)

func move(x, y int32) (int32, int32) {
	currentPos := position{}
	getCursorPosInfo.Call(uintptr(unsafe.Pointer(&currentPos)))
	setCursorPosInfo.Call(
		uintptr(currentPos.x+x),
		uintptr(currentPos.y+y),
	)

	return currentPos.x + x, currentPos.y + y
}

const (
	mouseEventFLeftDown  uint32 = 0x02
	mouseEventFLeftUp    uint32 = 0x04
	mouseEventFRightDown uint32 = 0x08
	mouseEventFRightUp   uint32 = 0x10
)

func click(t int) {
	switch t {
	case clickTypeLeft:
		mouseEvent(mouseEventFLeftDown | mouseEventFLeftUp)

	case clickTypeRight:
		mouseEvent(mouseEventFRightDown | mouseEventFRightUp)
	}
}

func mouseEvent(mouseEvent uint32) {
	mouseEventInfo.Call(
		uintptr(mouseEvent),
		uintptr(uint(0)),
		uintptr(uint(0)),
		uintptr(uint(0)),
		uintptr(uint(0)),
	)
}
