|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Unit | -  |
| --- | --- |
|  Visibility | Expert  |
|  Values | Send Receive  |

Selects which one of the send or receive features to control.

Possible values are:

- Send: Select the send information.
- Receive: Select the receive information.

### 27.7.18 CxpConnectionTestPacketCount

|  Name | CxpConnectionTestPacketCount[CxpConnectionSelector][CxpSendReceiveSelector]  |
| --- | --- |
|  Category | CoaXPress  |
|  Level | Recommended  |
|  Interface | Integer  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Reports the current count for the test packets on the connection selected by CxpConnectionSelector.

The CxpSendReceiveSelector can be used to choose between sent or received information.

Note that if this selector is omitted, CxpConnectionTestPacketCount reports the current count for test packets received by the device.

The transmission of these test packets is enabled by the CxpConnectionTestMode feature of the Device on the other end of the connection under test.

This feature can be read at any time while a test is running. When the test is not running, zero can be written to reset the counter between tests.

### 27.7.19 CxpErrorCounterSelector

|  Name | CxpErrorCounterSelector  |
| --- | --- |