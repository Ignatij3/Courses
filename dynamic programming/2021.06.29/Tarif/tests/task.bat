
@echo off
setlocal enabledelayedexpansion
for /L %%i in (1, 1, 25) do (
	if %%i LSS 10 (
		set in=00%%i.dat
		set out=00%%i.ans
	) else (
		set in=0%%i.dat
		set out=0%%i.ans
	)

	time <%0
	tarif.exe <!in! >tarif.res
	time <%0

	fc tarif.res !out!
	if errorlevel 1 (
		echo error on test #%%i
	)
)
