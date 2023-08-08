@echo off
cls
del input.txt
copy Input%1.txt input.txt 
corridor.exe
fc output.txt Answer%1.txt
pause