@echo off
del RemoteControl.exe
@echo on
go build
@echo off
if %ERRORLEVEL% EQU 0 (
  echo "start"
  RemoteControl.exe
) else (
  echo "errors"
  pause
)
