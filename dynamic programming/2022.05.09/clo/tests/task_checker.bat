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

    echo Checking %%i
	clo.exe <!in! >clo.res
	checker.exe <clo.res
)
del clo.res