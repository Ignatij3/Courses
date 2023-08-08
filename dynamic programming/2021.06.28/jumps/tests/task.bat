
@echo off
setlocal enabledelayedexpansion
for /L %%i in (0, 1, 10) do (
	if %%i LSS 10 (
		set in=jump.0%%i.in
		set out=jump.0%%i.sol
	) else (
		set in=jump.%%i.in
		set out=jump.%%i.sol
	)
	
	rem time <%0
	jump.exe <!in! >jumps.res
	rem time <%0
	
	fc jumps.res !out!
	del jumps.res
)
