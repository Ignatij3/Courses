@echo off
setlocal enabledelayedexpansion
copy /y nul clo.res
for /L %%i in (0, 1, 18) do (
	if %%i LSS 10 (
		set in=clo%%i.in
		set out=clo%%i.out
	) else (
		set in=clo%%i.in
		set out=clo%%i.out
	)

    echo running test #%%i

	echo !time!
	clo.exe <!in! >clo.res
	echo !time!

	fc clo.res !out! >nil
    if errorlevel 1 (
		echo error on test #%%i
	) else (
        echo test passed #%%i
    )
    echo:
)
del clo.res
