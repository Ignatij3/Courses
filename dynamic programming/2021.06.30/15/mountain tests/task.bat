
@echo off
setlocal enabledelayedexpansion
for /L %%i in (1, 1, 60) do (
	if %%i LSS 10 (
		set in=0%%i
		set out=0%%i.a
	) else (
		set in=%%i
		set out=%%i.a
	)

	time <%0
	mountain.exe <!in! >mountain.res
	time <%0

	fc mountain.res !out!
	if errorlevel 1 (
		echo error on test #%%i
	)
)
