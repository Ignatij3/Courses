@echo off
setlocal enabledelayedexpansion
copy /y nul cutting.res
for /L %%i in (1, 1, 21) do (
	if %%i LSS 10 (
		set in=0%%i
		set out=0%%i.a
	) else (
		set in=%%i
		set out=%%i.a
	)

    echo running test #%%i

	echo !time!
	cutting.exe <!in! >cutting.res
	echo !time!

	fc cutting.res !out! >nul
	if errorlevel 1 (
		echo error on test #%%i
	) else (
        echo test passed #%%i
    )
    echo:
)
del cutting.res
