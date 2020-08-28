@echo off
echo "made by Keen."

echo "run: %0"

set cur_dir=%~dp0
set bin_dir=%cur_dir%bin

go version 2>nul 1>nul
if %errorlevel% neq 0 (
	echo "golang not installed!"
	goto failed
)

go mod vendor

go build -mod=vendor -o %bin_dir%\WeatherNotify.exe %cur_dir%src\weather\main.go
go build -mod=vendor -o %bin_dir%\QuerySoftDir.exe %cur_dir%src\query_soft_dir\main.go

:successed
echo ">>>>>>>>>>>>>>>> run %0 successed!"
goto finished

:failed
echo ">>>>>>>>>>>>>>>> run %0 failed!"

:finished

pause