
@echo off
setlocal enabledelayedexpansion
for /L %%i in (1, 1, 3) do (
	if %%i LSS 10 (
		set in=0%%i.in
		set out=0%%i.out
	) else (
		set in=%%i.in
		set out=%%i.out
	)

	time <%0
	napoleon.exe <!in! >napoleon.res
	time <%0

	fc napoleon.res !out!
	if errorlevel 1 (
		echo error on test #%%i
	)
)
