/*
 provides the zLog programming interface to the Go language.
 Copyright (C) 2020 JA1ZLO.
*/
package main

import (
	"unsafe"
	"github.com/gonutz/w32"
)

func init() {
	OnLaunchEvent = onLaunchEvent
	OnWindowEvent = onWindowEvent
}

func onLaunchEvent() {
	h := w32.HMENU(GetUI("MainForm.MainMenu"))
	w32.AppendMenu(h, w32.MF_STRING, 810, "GO!")
	w32.DrawMenuBar(w32.HWND(GetUI("MainForm")))
}

func onWindowEvent(event uintptr) {
	msg := (*w32.MSG)(unsafe.Pointer(event))
	if msg.Message == w32.WM_COMMAND {
		if msg.WParam == 810 {
			DisplayToast("SCOOOP!!!")
		}
	}
}
