
@echo off
setlocal enabledelayedexpansion
for /L %%i in (1, 1, 20) do (
	if %%i LSS 10 (
		set in=coingame.0%%i.in
		set out=coingame.0%%i.out
	) else (
		set in=coingame.%%i.in
		set out=coingame.%%i.out
	)
	

	time <%0
	coingame.exe <!in! >coingame.res
	time <%0

	fc coingame.res !out!
	if errorlevel 1 (
		echo error on test #%%i
	)
	del coingame.res
)
