package main

import (
	"bytes"
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

const (
	ModAlt = 1 << iota
	ModCtrl
	ModShift
	ModWin
)

type Hotkey struct {
	Id        int // Unique id
	Modifiers int // Mask of modifiers
	KeyCode   int // Key code, e.g. 'A'
}
type MSG struct {
	HWND   uintptr
	UINT   uintptr
	WPARAM int16
	LPARAM int64
	DWORD  int32
	POINT  struct{ X, Y int64 }
}

// String returns a human-friendly display name of the hotkey
// such as "Hotkey[Id: 1, Alt+Ctrl+O]"
func (h *Hotkey) String() string {
	mod := &bytes.Buffer{}
	if h.Modifiers&ModAlt != 0 {
		mod.WriteString("Alt+")
	}
	if h.Modifiers&ModCtrl != 0 {
		mod.WriteString("Ctrl+")
	}
	if h.Modifiers&ModShift != 0 {
		mod.WriteString("Shift+")
	}
	if h.Modifiers&ModWin != 0 {
		mod.WriteString("Win+")
	}
	return fmt.Sprintf("Hotkey[Id: %d, %s%c]", h.Id, mod, h.KeyCode)
}

func main() {
	user32 := syscall.MustLoadDLL("user32")
	defer user32.Release()
	reghotkey := user32.MustFindProc("RegisterHotKey")

	// Hotkeys to listen to:
	// https://msdn.microsoft.com/en-us/library/windows/desktop/dd375731(v=vs.85).aspx
	keys := map[int16]*Hotkey{
		1: &Hotkey{1, ModCtrl + ModAlt, 'Y'},
		2: &Hotkey{2, ModWin + ModAlt, 'Y'},
	}

	// Register hotkeys:
	for _, v := range keys {
		// https://msdn.microsoft.com/en-us/library/windows/desktop/ms646309(v=vs.85).aspx
		r1, _, err := reghotkey.Call(
			0, uintptr(v.Id), uintptr(v.Modifiers), uintptr(v.KeyCode))
		if r1 == 1 {
			fmt.Println("Registered", v)
		} else {
			fmt.Println("Failed to register", v, ", error:", err)
		}
	}

	peekmsg := user32.MustFindProc("PeekMessageW")
	for {
		var msg = &MSG{}
		peekmsg.Call(uintptr(unsafe.Pointer(msg)), 0, 0, 0, 1)

		// Registered id is in the WPARAM field:
		if id := msg.WPARAM; id != 0 {
			fmt.Println("Hotkey pressed:", keys[id])
			if id == 1 { // Win+F4 = Exit
				fmt.Println(keys[id], "pressed, goodbye...")
				return
			}
		}

		time.Sleep(time.Millisecond * 50)
	}
}
