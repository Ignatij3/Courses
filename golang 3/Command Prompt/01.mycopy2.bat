@echo off
if "%1"=="" (
  echo ������� �ந����쭮� ������⢮
  echo 䠩��� � 㪠����� ��⠫��
  echo �맮�:  
  echo  %~NX0 directory file1 file2 ...
  exit /B
)
:getfile
if not "%2"=="" (
  copy /B %2 %1 >nul
  shift /2 
  goto getfile 
)
echo ����஢���� �����祭�
