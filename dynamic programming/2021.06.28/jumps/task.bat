@echo off
for /L %%i in (0, 1, 9) do (
	solution.exe <jump.0%%i.in >jumps.res
	fc jumps.res jump.0%%i.sol
	if errorlevel 1 (
		echo error on test #%%i
		exit /b
	)
)    
solution.exe <jump.10.in >jumps.res
fc jumps.res jump.10.sol
if errorlevel 1 (
	echo error on test #10
)
pause
