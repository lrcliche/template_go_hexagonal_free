@echo off
setlocal enabledelayedexpansion

if "%~1"=="" (
  echo Usage: scripts\add_feature.bat ^<feature_name^>
  exit /b 1
)

go run ./tools/add_feature %1
if errorlevel 1 exit /b %errorlevel%

endlocal
