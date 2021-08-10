@echo off
if %1 lss %2 echo %1 ^< %2
if %1 gtr %2 echo %1 ^> %2
if %1 == %2 echo %1 ^= %2