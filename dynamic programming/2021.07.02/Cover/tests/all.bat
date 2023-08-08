@echo off
cls
for /L %%i in (1,1,9) do call alz.bat 0%%i
for /L %%i in (10,1,10) do call alz.bat %%i
