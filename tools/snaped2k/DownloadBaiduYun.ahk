#NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
; #Warn  ; Enable warnings to assist with detecting common errors.
SendMode Input  ; Recommended for new scripts due to its superior speed and reliability.
SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.

^j::
	Loop, Read, list.txt
	{
		WinActivate, 欢迎使用百度网盘
		Click 567, 118
		Sleep, 300
		Click 10886, 10254
		WinWaitActive, 新建离线下载任务窗口
		Click 58, 103
		Send, %A_LoopReadLine%
		Click 477, 253
		WinWaitActive, 离线下载任务列表
		Click 483, 363
	}
	Return

::qq::
	MsgBox You typed "qq".
	Return
