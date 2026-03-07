@echo off
setlocal EnableDelayedExpansion
cd /d %~dp0\..

echo.
echo ==^> Hexagonal architecture guard

findstr /S /N /I "infrastructure/" domain\*.go >nul 2>&1
if %errorlevel%==0 (
  echo [FAIL] domain must not depend on infrastructure
  exit /b 1
) else (
  echo [PASS] domain does not import infrastructure
)

findstr /S /N /I "pgxpool sql.DB gorm bun." presentation\*.go >nul 2>&1
if %errorlevel%==0 (
  echo [FAIL] presentation must not access database implementations
  exit /b 1
) else (
  echo [PASS] presentation does not access database implementations
)

findstr /S /N /I "infrastructure/" application\*.go >nul 2>&1
if %errorlevel%==0 (
  echo [FAIL] application must not depend on infrastructure
  exit /b 1
) else (
  echo [PASS] application does not import infrastructure
)

findstr /S /N /I "domain/ports" application\services\*.go >nul 2>&1
if not %errorlevel%==0 (
  echo [FAIL] application services must depend on repository ports
  exit /b 1
) else (
  echo [PASS] application services depend on repository ports
)

go test ./...
if not %errorlevel%==0 (
  echo [FAIL] go test ./...
  exit /b 1
)

echo [PASS] go test ./...
echo Architecture checks completed successfully.
