|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- Off: Test mode is disabled.
- Mode1: Test mode is one.

This can be used to test communication errors of the system cabling between devices.

When enabled, this feature results in special test packets being sent continuously by the Device on the connection specified by CxpConnectionSelector.

The Device receiving the test packet on the other end of the connection can check for errors by reading its own corresponding CxpConnectionTestErrorCount and CxpConnectionTestPacketCount features.

Typically, the test will need to be run for some time (e.g. minutes) to get a meaningful error rate.

### 27.7.16 CxpConnectionTestErrorCount

|  Name | CxpConnectionTestErrorCount[CxpConnectionSelector]  |
| --- | --- |
|  Category | CoaXPress  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Reports the current connection error count for test packets received by the device on the connection selected by CxpConnectionSelector.

The transmission of those test packets is enabled by the CxpConnectionTestMode feature of the Device on the other end of the connection under test.

This feature can be read at any time while a test is running. It can be written to zero when a test is not running to reset the counter between tests.

### 27.7.17 CxpSendReceiveSelector

|  Name | CxpSendReceiveSelector  |
| --- | --- |
|  Category | CoaXPress  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |