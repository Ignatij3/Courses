@echo off
for /L %%n in (0, 1, 9) do (
	echo %%n > ..\Batch_results\0%%n.txt
)
for /L %%n in (10, 1, 99) do (
	echo %%n > ..\Batch_results\%%n.txt
)