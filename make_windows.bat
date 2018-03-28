@echo off

setlocal

if exist make_windows.bat goto ok
echo install.bat must be run from its folder
goto end

:ok

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0;%OLDGOPATH%

gofmt -w src
go clean
go install game_server
:end
echo finished
pause