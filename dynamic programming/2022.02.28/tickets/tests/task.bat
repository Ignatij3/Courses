@echo off
setlocal enabledelayedexpansion
copy /y nul tickets.res
for /L %%i in (1, 1, 10) do (
	if %%i LSS 10 (
		set in=%%i
		set out=%%i.a
	) else (
		set in=%%i
		set out=%%i.a
	)

	echo %time%
	tickets.exe <!in! >tickets.res
	echo %time%

	fc tickets.res !out!
	if errorlevel 1 (
		echo error on test #%%i
	)
)
del tickets.res
