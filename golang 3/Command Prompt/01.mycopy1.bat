@echo off
set todir=%1
if "%todir%"=="" goto intro
:getfile
shift
if "%1"=="" goto end
copy /B %1 %todir% >nul
goto getfile
:intro
echo ������� �ந����쭮� ������⢮
echo 䠩��� � 㪠����� ��⠫��
echo �맮�:  
echo   %0 directory file1 file2 ...
exit /B
:end
echo ����஢���� �����祭�
