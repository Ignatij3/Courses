@echo off
setlocal enabledelayedexpansion
copy /y nul alibi.res
for /L %%i in (1, 1, 31) do (
	if %%i LSS 10 (
		set in=0%%i
		set out=0%%i.a
	) else (
		set in=%%i
		set out=%%i.a
	)

    echo running test #%%i

	echo !time!
	alibi.exe <!in! >alibi.res
	echo !time!

	fc alibi.res !out! >nul
	if errorlevel 1 (
		echo error on test #%%i
        	goto :eof
	) else (
        	echo test passed #%%i
    	)
    echo:
)
del alibi.res
