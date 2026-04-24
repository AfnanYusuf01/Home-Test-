###############################################################################
# Test Scenarios — Paket Data REST API
# Full comprehensive test: success & failure cases
# Uses temp file approach to avoid PowerShell JSON escaping issues with curl.exe
###############################################################################

$BASE = "http://localhost:3000/api"
$PASS = 0
$FAIL = 0
$TMP = "$PSScriptRoot\tmp_test_body.json"

function Send-Request {
    param(
        [string]$Method,
        [string]$Url,
        [string]$Body
    )

    if ($Body) {
        [System.IO.File]::WriteAllText($TMP, $Body)
        $raw = curl.exe -s -w "`n%{http_code}" -X $Method -H "Content-Type: application/json" -d "@$TMP" $Url 2>$null
    } else {
        $raw = curl.exe -s -w "`n%{http_code}" -X $Method $Url 2>$null
    }

    $lines = ($raw -split "`n") | Where-Object { $_ -ne "" }
    $statusCode = [int]$lines[-1]
    $responseBody = ($lines[0..($lines.Length-2)]) -join ""
    $content = $responseBody | ConvertFrom-Json -ErrorAction SilentlyContinue

    return @{
        StatusCode = $statusCode
        Body = $content
        Raw = $responseBody
    }
}

function Test-Endpoint {
    param(
        [string]$Name,
        [string]$Method,
        [string]$Url,
        [string]$Body,
        [int]$ExpectedStatus,
        [bool]$ExpectSuccess
    )

    Write-Host ""
    Write-Host "--- TEST: $Name ---" -ForegroundColor Cyan
    Write-Host "  $Method $Url" -ForegroundColor DarkGray
    if ($Body) { Write-Host "  Body: $Body" -ForegroundColor DarkGray }

    $r = Send-Request -Method $Method -Url $Url -Body $Body

    $statusOk = ($r.StatusCode -eq $ExpectedStatus)
    $successOk = ($r.Body.success -eq $ExpectSuccess)

    if ($statusOk -and $successOk) {
        Write-Host "  PASS | Status: $($r.StatusCode) | $($r.Body.message)" -ForegroundColor Green
        $script:PASS++
    } else {
        Write-Host "  FAIL | Expected: $ExpectedStatus/success=$ExpectSuccess | Got: $($r.StatusCode)/success=$($r.Body.success)" -ForegroundColor Red
        Write-Host "  Response: $($r.Raw)" -ForegroundColor Yellow
        $script:FAIL++
    }

    return $r.Body
}

# ============================================================================
Write-Host ""
Write-Host "================================================================" -ForegroundColor Magenta
Write-Host "     PAKET DATA REST API - COMPREHENSIVE TEST SCENARIOS" -ForegroundColor Magenta
Write-Host "================================================================" -ForegroundColor Magenta

# ============================================================================
# 1. WELCOME
# ============================================================================
Write-Host "`n>> SECTION 1: Welcome Endpoint <<" -ForegroundColor Yellow

Test-Endpoint -Name "GET / Welcome" `
    -Method GET -Url "http://localhost:3000/" `
    -ExpectedStatus 200 -ExpectSuccess $true

# ============================================================================
# 2. USER CRUD - SUCCESS
# ============================================================================
Write-Host "`n>> SECTION 2: User CRUD <<" -ForegroundColor Yellow

$r = Test-Endpoint -Name "POST Create User 1" `
    -Method POST -Url "$BASE/users" `
    -Body '{"name":"Afnan Yusuf","phone_number":"08123456789"}' `
    -ExpectedStatus 201 -ExpectSuccess $true
$U1 = if ($r -and $r.data) { $r.data.id } else { 1 }

$r = Test-Endpoint -Name "POST Create User 2" `
    -Method POST -Url "$BASE/users" `
    -Body '{"name":"Budi Santoso","phone_number":"08987654321"}' `
    -ExpectedStatus 201 -ExpectSuccess $true
$U2 = if ($r -and $r.data) { $r.data.id } else { 2 }

Test-Endpoint -Name "GET All Users" `
    -Method GET -Url "$BASE/users" `
    -ExpectedStatus 200 -ExpectSuccess $true

Test-Endpoint -Name "GET User by ID ($U1)" `
    -Method GET -Url "$BASE/users/$U1" `
    -ExpectedStatus 200 -ExpectSuccess $true

Test-Endpoint -Name "PUT Update User ($U1)" `
    -Method PUT -Url "$BASE/users/$U1" `
    -Body '{"name":"Afnan Updated"}' `
    -ExpectedStatus 200 -ExpectSuccess $true

# USER - FAILURE
Test-Endpoint -Name "POST User - Empty Body (validation)" `
    -Method POST -Url "$BASE/users" `
    -Body '{}' `
    -ExpectedStatus 400 -ExpectSuccess $false

Test-Endpoint -Name "POST User - Name Too Short" `
    -Method POST -Url "$BASE/users" `
    -Body '{"name":"AB","phone_number":"08123456789"}' `
    -ExpectedStatus 400 -ExpectSuccess $false

Test-Endpoint -Name "POST User - Phone Too Short" `
    -Method POST -Url "$BASE/users" `
    -Body '{"name":"Test User","phone_number":"0812"}' `
    -ExpectedStatus 400 -ExpectSuccess $false

Test-Endpoint -Name "POST User - Duplicate Phone (409)" `
    -Method POST -Url "$BASE/users" `
    -Body '{"name":"Dup User","phone_number":"08123456789"}' `
    -ExpectedStatus 409 -ExpectSuccess $false

Test-Endpoint -Name "GET User - Not Found (999)" `
    -Method GET -Url "$BASE/users/999" `
    -ExpectedStatus 404 -ExpectSuccess $false

Test-Endpoint -Name "GET User - Invalid ID (abc)" `
    -Method GET -Url "$BASE/users/abc" `
    -ExpectedStatus 400 -ExpectSuccess $false

Test-Endpoint -Name "PUT User - Not Found (999)" `
    -Method PUT -Url "$BASE/users/999" `
    -Body '{"name":"Ghost"}' `
    -ExpectedStatus 404 -ExpectSuccess $false

Test-Endpoint -Name "POST User - Invalid JSON" `
    -Method POST -Url "$BASE/users" `
    -Body 'not-json' `
    -ExpectedStatus 400 -ExpectSuccess $false

# ============================================================================
# 3. PAKET DATA CRUD
# ============================================================================
Write-Host "`n>> SECTION 3: Paket Data CRUD <<" -ForegroundColor Yellow

$r = Test-Endpoint -Name "POST Create Paket Data 1" `
    -Method POST -Url "$BASE/paket-data" `
    -Body '{"name":"Paket Internet 10GB","price":50000,"quota":10,"active_period":30}' `
    -ExpectedStatus 201 -ExpectSuccess $true
$P1 = if ($r -and $r.data) { $r.data.id } else { 1 }

$r = Test-Endpoint -Name "POST Create Paket Data 2" `
    -Method POST -Url "$BASE/paket-data" `
    -Body '{"name":"Paket Internet 25GB","price":100000,"quota":25,"active_period":30}' `
    -ExpectedStatus 201 -ExpectSuccess $true
$P2 = if ($r -and $r.data) { $r.data.id } else { 2 }

Test-Endpoint -Name "GET All Paket Data" `
    -Method GET -Url "$BASE/paket-data" `
    -ExpectedStatus 200 -ExpectSuccess $true

Test-Endpoint -Name "GET Paket Data by ID ($P1)" `
    -Method GET -Url "$BASE/paket-data/$P1" `
    -ExpectedStatus 200 -ExpectSuccess $true

Test-Endpoint -Name "PUT Update Paket Data ($P1)" `
    -Method PUT -Url "$BASE/paket-data/$P1" `
    -Body '{"name":"Paket Internet 15GB","quota":15,"price":65000}' `
    -ExpectedStatus 200 -ExpectSuccess $true

# PAKET DATA - FAILURE
Test-Endpoint -Name "POST Paket Data - Empty Body" `
    -Method POST -Url "$BASE/paket-data" `
    -Body '{}' `
    -ExpectedStatus 400 -ExpectSuccess $false

Test-Endpoint -Name "POST Paket Data - Negative Price" `
    -Method POST -Url "$BASE/paket-data" `
    -Body '{"name":"Bad Paket","price":-1000,"quota":10,"active_period":30}' `
    -ExpectedStatus 400 -ExpectSuccess $false

Test-Endpoint -Name "POST Paket Data - Zero Quota" `
    -Method POST -Url "$BASE/paket-data" `
    -Body '{"name":"Zero Paket","price":50000,"quota":0,"active_period":30}' `
    -ExpectedStatus 400 -ExpectSuccess $false

Test-Endpoint -Name "GET Paket Data - Not Found (999)" `
    -Method GET -Url "$BASE/paket-data/999" `
    -ExpectedStatus 404 -ExpectSuccess $false

Test-Endpoint -Name "PUT Paket Data - Not Found (999)" `
    -Method PUT -Url "$BASE/paket-data/999" `
    -Body '{"name":"Ghost Paket"}' `
    -ExpectedStatus 404 -ExpectSuccess $false

# ============================================================================
# 4. TRANSAKSI
# ============================================================================
Write-Host "`n>> SECTION 4: Transaksi <<" -ForegroundColor Yellow

$r = Test-Endpoint -Name "POST Transaksi - User$U1 beli Paket$P1" `
    -Method POST -Url "$BASE/transaksi" `
    -Body "{""user_id"":$U1,""paket_data_id"":$P1}" `
    -ExpectedStatus 201 -ExpectSuccess $true
$T1 = if ($r -and $r.data) { $r.data.id } else { 1 }

$r = Test-Endpoint -Name "POST Transaksi - User$U2 beli Paket$P2" `
    -Method POST -Url "$BASE/transaksi" `
    -Body "{""user_id"":$U2,""paket_data_id"":$P2}" `
    -ExpectedStatus 201 -ExpectSuccess $true
$T2 = if ($r -and $r.data) { $r.data.id } else { 2 }

Test-Endpoint -Name "POST Transaksi - User$U1 beli Paket$P2 (multi)" `
    -Method POST -Url "$BASE/transaksi" `
    -Body "{""user_id"":$U1,""paket_data_id"":$P2}" `
    -ExpectedStatus 201 -ExpectSuccess $true

Test-Endpoint -Name "GET All Transaksi" `
    -Method GET -Url "$BASE/transaksi" `
    -ExpectedStatus 200 -ExpectSuccess $true

Test-Endpoint -Name "GET Transaksi by ID ($T1)" `
    -Method GET -Url "$BASE/transaksi/$T1" `
    -ExpectedStatus 200 -ExpectSuccess $true

# TRANSAKSI - FAILURE
Test-Endpoint -Name "POST Transaksi - Empty Body" `
    -Method POST -Url "$BASE/transaksi" `
    -Body '{}' `
    -ExpectedStatus 400 -ExpectSuccess $false

Test-Endpoint -Name "POST Transaksi - User Not Found" `
    -Method POST -Url "$BASE/transaksi" `
    -Body "{""user_id"":999,""paket_data_id"":$P1}" `
    -ExpectedStatus 404 -ExpectSuccess $false

Test-Endpoint -Name "POST Transaksi - Paket Not Found" `
    -Method POST -Url "$BASE/transaksi" `
    -Body "{""user_id"":$U1,""paket_data_id"":999}" `
    -ExpectedStatus 404 -ExpectSuccess $false

Test-Endpoint -Name "GET Transaksi - Not Found (999)" `
    -Method GET -Url "$BASE/transaksi/999" `
    -ExpectedStatus 404 -ExpectSuccess $false

# ============================================================================
# 5. SOFT DELETE & HISTORICAL INTEGRITY
# ============================================================================
Write-Host "`n>> SECTION 5: Soft Delete & Historical Integrity <<" -ForegroundColor Yellow

# Soft delete User1
Test-Endpoint -Name "DELETE User$U1 (soft delete)" `
    -Method DELETE -Url "$BASE/users/$U1" `
    -ExpectedStatus 200 -ExpectSuccess $true

# User1 tidak bisa diakses lagi
Test-Endpoint -Name "GET User$U1 after delete - 404" `
    -Method GET -Url "$BASE/users/$U1" `
    -ExpectedStatus 404 -ExpectSuccess $false

# Delete lagi - sudah tidak ada
Test-Endpoint -Name "DELETE User$U1 again - already deleted 404" `
    -Method DELETE -Url "$BASE/users/$U1" `
    -ExpectedStatus 404 -ExpectSuccess $false

# KUNCI: Transaksi User1 MASIH ADA (historical integrity)
Test-Endpoint -Name "GET Transaksi$T1 after user delete - STILL EXISTS" `
    -Method GET -Url "$BASE/transaksi/$T1" `
    -ExpectedStatus 200 -ExpectSuccess $true

# Tidak bisa buat transaksi baru dengan user yang sudah dihapus
Test-Endpoint -Name "POST Transaksi with deleted user - BLOCKED 404" `
    -Method POST -Url "$BASE/transaksi" `
    -Body "{""user_id"":$U1,""paket_data_id"":$P2}" `
    -ExpectedStatus 404 -ExpectSuccess $false

# Soft delete Paket Data 1
Test-Endpoint -Name "DELETE PaketData$P1 (soft delete)" `
    -Method DELETE -Url "$BASE/paket-data/$P1" `
    -ExpectedStatus 200 -ExpectSuccess $true

# Paket Data 1 tidak bisa diakses
Test-Endpoint -Name "GET PaketData$P1 after delete - 404" `
    -Method GET -Url "$BASE/paket-data/$P1" `
    -ExpectedStatus 404 -ExpectSuccess $false

# KUNCI: Transaksi dengan PaketData1 MASIH ADA
Test-Endpoint -Name "GET Transaksi$T1 after paket delete - STILL EXISTS" `
    -Method GET -Url "$BASE/transaksi/$T1" `
    -ExpectedStatus 200 -ExpectSuccess $true

# Tidak bisa buat transaksi dengan paket yang sudah dihapus
Test-Endpoint -Name "POST Transaksi with deleted paket - BLOCKED 404" `
    -Method POST -Url "$BASE/transaksi" `
    -Body "{""user_id"":$U2,""paket_data_id"":$P1}" `
    -ExpectedStatus 404 -ExpectSuccess $false

# Semua transaksi masih ada
Test-Endpoint -Name "GET All Transaksi - all 3 still exist" `
    -Method GET -Url "$BASE/transaksi" `
    -ExpectedStatus 200 -ExpectSuccess $true

# ============================================================================
# 6. EDGE CASES
# ============================================================================
Write-Host "`n>> SECTION 6: Edge Cases <<" -ForegroundColor Yellow

Test-Endpoint -Name "DELETE User 999 - Not Found" `
    -Method DELETE -Url "$BASE/users/999" `
    -ExpectedStatus 404 -ExpectSuccess $false

Test-Endpoint -Name "DELETE Paket 999 - Not Found" `
    -Method DELETE -Url "$BASE/paket-data/999" `
    -ExpectedStatus 404 -ExpectSuccess $false

Test-Endpoint -Name "GET /api/nonexistent - 404" `
    -Method GET -Url "$BASE/nonexistent" `
    -ExpectedStatus 404 -ExpectSuccess $false

Test-Endpoint -Name "PUT User with broken JSON" `
    -Method PUT -Url "$BASE/users/$U2" `
    -Body 'broken{json' `
    -ExpectedStatus 400 -ExpectSuccess $false

# ============================================================================
# CLEANUP
# ============================================================================
if (Test-Path $TMP) { Remove-Item $TMP -Force }

# ============================================================================
# SUMMARY
# ============================================================================
Write-Host ""
Write-Host "================================================================" -ForegroundColor Magenta
Write-Host "                     TEST RESULTS SUMMARY" -ForegroundColor Magenta
Write-Host "================================================================" -ForegroundColor Magenta
Write-Host ""
Write-Host "  PASSED: $PASS" -ForegroundColor Green
Write-Host "  FAILED: $FAIL" -ForegroundColor Red
Write-Host "  TOTAL:  $($PASS + $FAIL)" -ForegroundColor White
Write-Host ""
if ($FAIL -eq 0) {
    Write-Host "  >>> ALL TESTS PASSED! <<<" -ForegroundColor Green
} else {
    Write-Host "  >>> $FAIL TEST(S) FAILED <<<" -ForegroundColor Red
}
Write-Host ""
