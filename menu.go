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

func init() {
	OnAttachEvent = onLaunchEvent
}

func onLaunchEvent(rule, name string) {
	hMenu := win.HMENU(GetUI("MainForm.MainMenu"))
	var mii win.MENUITEMINFO
	mii.CbSize = uint32(unsafe.Sizeof(mii))
	mii.FMask = win.MIIM_TYPE
	mii.FType = win.MFT_STRING
	mii.DwTypeData = syscall.StringToUTF16Ptr("MENU1")
	mii.Cch = uint32(len([]rune("MENU!")))
	if !win.InsertMenuItem(hMenu, uint32(4), true, &mii) {
		DisplayModal("failed")
	}
	mii.DwTypeData = syscall.StringToUTF16Ptr("MENU2")
	if !win.InsertMenuItem(hMenu, uint32(5), true, &mii) {
		DisplayModal("failed")
	}
	if !win.DrawMenuBar(win.HWND(GetUI("MainForm"))) {
		DisplayModal("failed")
	}
}
