@echo off
if "%1"=="" (
  echo Копирует произвольное количество
  echo файлов в указанный каталог
  echo Вызов:  
  echo  %~NX0 directory file1 file2 ...
  exit /B
)
:getfile
if not "%2"=="" (
  copy /B %2 %1 >nul
  shift /2 
  goto getfile 
)
echo Копирование закончено
