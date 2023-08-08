
@echo off
setlocal enabledelayedexpansion
for /L %%i in (0, 1, 10) do (
	if %%i LSS 10 (
		set in=rec.0%%i.in
		set out=rec.0%%i.sol
	) else (
		set in=rec.%%i.in
		set out=rec.%%i.sol
	)
	

	time <%0
	rectangles.exe <!in! >rectangles.res
	time <%0

	fc rectangles.res !out!
	if errorlevel 1 (
		echo error on test #%%i
	)
	del rectangles.res
)
