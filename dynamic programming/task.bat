@echo off
setlocal enabledelayedexpansion
copy /y nul EXECUTABLE_NAME.res
for /L %%i in (0, 1, 10) do (
	if %%i LSS 10 (
		set in=0%%i.in
		set out=0%%i.out
	) else (
		set in=%%i.in
		set out=%%i.out
	)

    	echo running test #%%i

	echo !time!
	EXECUTABLE_NAME.exe <!in! >EXECUTABLE_NAME.res
	echo !time!

	fc EXECUTABLE_NAME.res !out! >nul
	if errorlevel 1 (
		echo error on test #%%i
		goto :eof
	) else (
        	echo test passed #%%i
    	)
    echo:
)
del EXECUTABLE_NAME.res
