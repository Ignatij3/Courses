@echo off
for /L %%l in (1, 1, 20) do (
	for /F "delims=," %%n in (file.dat) do (
		if %%l lss 10 (
			echo %%n > file.0%%l.dat
		) else (
			echo %%n > file.%%l.dat
		)
	)
)
echo Done!