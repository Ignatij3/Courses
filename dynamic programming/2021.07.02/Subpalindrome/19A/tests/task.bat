
@echo off
setlocal enabledelayedexpansion
for /L %%i in (1, 1, 2) do (
	set in=input0%%i.dat
	set out=output0%%i.dat

	time <%0
	subpalindrome.exe <!in! >subpalindrome.res
	time <%0

	fc subpalindrome.res !out!
	if errorlevel 1 (
		echo error on test #%%i
	)
	del subpalindrome.res
)
