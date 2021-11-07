/*
 provides the zLog programming interface to the Go language.
 Copyright (C) 2020 JA1ZLO.
*/
package main

import (
	"syscall"
	"unsafe"
	"github.com/lxn/win"
)

const MENU_ID = 114514

func init() {
	OnLaunchEvent = onLaunchEvent
	OnHandleEvent = onHandleEvent
}

func onHandleEvent(msg uintptr) {
	message := (*win.MSG)(unsafe.Pointer(msg))
	if message.Message == win.WM_COMMAND {
		DisplayToast("CLICKED! %d", message.WParam)
	}
}

func onLaunchEvent() {
	hMenu := win.HMENU(GetUI("MainForm.MainMenu"))
	title := "MENU"
	var mii win.MENUITEMINFO
	mii.CbSize = uint32(unsafe.Sizeof(mii))
	mii.FMask = win.MIIM_TYPE | win.MIIM_ID
	mii.FType = win.MFT_STRING
	mii.DwTypeData = syscall.StringToUTF16Ptr(title)
	mii.Cch = uint32(len([]rune(title)))
	mii.WID = MENU_ID
	win.InsertMenuItem(hMenu, uint32(4), true, &mii)
	win.DrawMenuBar(win.HWND(GetUI("MainForm")))
}
