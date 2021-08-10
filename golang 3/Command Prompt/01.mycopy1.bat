@echo off
set todir=%1
if "%todir%"=="" goto intro
:getfile
shift
if "%1"=="" goto end
copy /B %1 %todir% >nul
goto getfile
:intro
echo Копирует произвольное количество
echo файлов в указанный каталог
echo Вызов:  
echo   %0 directory file1 file2 ...
exit /B
:end
echo Копирование закончено
