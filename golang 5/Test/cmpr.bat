
@echo off
setlocal enabledelayedexpansion
for /L %%n in (1, 1, 26) do (
  if %%n LSS 10 (
    set in=input0%%n.txt
    set out=answer0%%n.txt
  ) else (
    set in=input%%n.txt
    set out=answer%%n.txt
  )

  time <%0
  Trains.exe <!in! >train.res
  time <%0
  
  fc train.res !out!
  if errorlevel 1  (
    type !out!
    type train.res 
  )
  del train.res
)