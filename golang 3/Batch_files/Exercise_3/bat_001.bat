@echo off
for /L %%n in (0, 1, 99) do (
	if %%n lss 10 (
		echo 0%%n > 0%%n.txt
	) else (
		echo %%n > %%n.txt
	)
)
echo Your files will be deleted, press ctrl + c, to stop
pause
for /L %%n in (0, 1, 99) do (
	if %%n lss 10 (
		del /Q 0%%n.txt
	) else (
		del /Q %%n.txt
	)
)