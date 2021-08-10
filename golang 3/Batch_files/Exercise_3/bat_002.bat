@echo off
for /L %%n in (0, 1, 99) do (
	if %%n lss 10 (
		for /L %%k in (0, 1, %%n - 1) do (
			echo 0%%n >> ..\Batch_results\0%%n.txt
		)
	) else (
		for /L %%k in (0, 1, %%n - 1) do (
			echo %%n >> ..\Batch_results\%%n.txt
		)
	)
)
echo Your files will be deleted, press ctrl + c, to stop
pause
for /L %%n in (0, 1, 99) do (
	if %%n lss 10 (
		del /Q ..\Batch_results\0%%n.txt
	) else (
		del /Q ..\Batch_results\%%n.txt
	)
)