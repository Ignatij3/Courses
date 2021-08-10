@echo off
:Start
rem choice /M "first second third fourth end" 12345
choice /C 12345
if errorlevel 5 (
	@echo.
	echo Gold ^= %errorlevel%
	goto End
)

if errorlevel 4 (
	@echo.
	echo Gold ^= %errorlevel%
	goto Point
)

if errorlevel 3 (
	@echo.
	echo Gold ^= %errorlevel%
	goto Point
)

if errorlevel 2 (
	@echo.
	echo Gold ^= %errorlevel%
	goto Point
)

if errorlevel 1 (
	@echo.
	echo Gold ^= %errorlevel%
	goto Point
)
:Point
goto Start
:End
echo End of programm