Error 1 
[error] failed to initialize database, got error failed to connect to `host=localhost user=postgres database=testdb`: failed SASL auth (FATAL: password authentication failed for user "postgres" (SQLSTATE 28P01))
--- FAIL: TestRegisterDevice (0.13s)
panic: Failed to connect to test database [recovered]
        panic: Failed to connect to test database

goroutine 18 [running]:
testing.tRunner.func1.2({0x75b460, 0x95f940})
        C:/Program Files/Go/src/testing/testing.go:1632 +0x225
testing.tRunner.func1()
        C:/Program Files/Go/src/testing/testing.go:1635 +0x359
panic({0x75b460?, 0x95f940?})
        C:/Program Files/Go/src/runtime/panic.go:785 +0x132
device-api.setupTestDB()
        c:/Users/priyajit.gantayat/Documents/Mini Project Backup - Copy - Copy/device-api/device_test.go:22 +0x175
device-api.setupTestRouter()
        c:/Users/priyajit.gantayat/Documents/Mini Project Backup - Copy - Copy/device-api/device_test.go:33 +0xf
device-api.TestRegisterDevice(0xc000136340)
        c:/Users/priyajit.gantayat/Documents/Mini Project Backup - Copy - Copy/device-api/device_test.go:40 +0x2a
testing.tRunner(0xc000136340, 0x8a5d30)
        C:/Program Files/Go/src/testing/testing.go:1690 +0xcb
created by testing.(*T).Run in goroutine 1
        C:/Program Files/Go/src/testing/testing.go:1743 +0x377
FAIL    device-api      0.233s

Error 2
[error] failed to initialize database, got error failed to connect to `host=localhost user=postgres database=testdb`: failed SASL auth (FATAL: password authentication failed for user "postgres" (SQLSTATE 28P01))
--- FAIL: TestListDevices (3.70s)
panic: Failed to connect to test database [recovered]
        panic: Failed to connect to test database

goroutine 7 [running]:
testing.tRunner.func1.2({0x11db460, 0x13df940})
        C:/Program Files/Go/src/testing/testing.go:1632 +0x225
testing.tRunner.func1()
        C:/Program Files/Go/src/testing/testing.go:1635 +0x359
panic({0x11db460?, 0x13df940?})
        C:/Program Files/Go/src/runtime/panic.go:785 +0x132
device-api.setupTestDB()
        c:/Users/priyajit.gantayat/Documents/Mini Project Backup - Copy - Copy/device-api/device_test.go:22 +0x175
device-api.setupTestRouter()
        c:/Users/priyajit.gantayat/Documents/Mini Project Backup - Copy - Copy/device-api/device_test.go:33 +0xf
device-api.TestListDevices(0xc0000c9040)
        c:/Users/priyajit.gantayat/Documents/Mini Project Backup - Copy - Copy/device-api/device_test.go:70 +0x1f
testing.tRunner(0xc0000c9040, 0x1325d28)
        C:/Program Files/Go/src/testing/testing.go:1690 +0xcb
created by testing.(*T).Run in goroutine 1
        C:/Program Files/Go/src/testing/testing.go:1743 +0x377
FAIL    device-api      3.793s

