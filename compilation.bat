@echo off
cls

cd %GOPATH%
if not exist "bin" mkdir bin
cd bin

go build main
pause
main

pause