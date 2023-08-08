
@echo off
setlocal enabledelayedexpansion
for /L %%i in (1, 1, 2) do (
	set in=input0%%i.dat
	set out=output0%%i.dat

	time <%0
	subpalindrome_with_str.exe <!in! >subpalindrome_with_str.res
	time <%0

	fc subpalindrome_with_str.res !out!
	if errorlevel 1 (
		echo error on test #%%i
	)
	del subpalindrome_with_str.res
)
