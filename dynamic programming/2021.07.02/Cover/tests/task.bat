
@echo off
setlocal enabledelayedexpansion
for /L %%i in (1, 1, 10) do (
	type NUL > corridor.res
	if %%i LSS 10 (
		set in=Input0%%i.txt
		set out=Answer0%%i.txt
	) else (
		set in=Input%%i.txt
		set out=Answer%%i.txt
	)

	time <%0
	corridor.exe <!in! >corridor.res
	time <%0

	fc corridor.res !out!
	if errorlevel 1 (
		echo error on test #%%i
	)
	del corridor.res
)
