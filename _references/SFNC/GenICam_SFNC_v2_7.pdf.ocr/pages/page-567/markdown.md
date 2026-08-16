|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

This feature provides the duplicate count in striped system. A non-zero value sets the number of duplicate images sent to sub-Devices.

### 27.7.14 CxpConnectionSelector

|  Name | CxpConnectionSelector  |
| --- | --- |
|  Category | CoaXPress  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Selects the CoaXPress physical connection to control.

Note that this selector should be set to 0, or omitted, when controlling features relating to the CoaXPress uplink from the Receiver Device to the Transmitter Device, because only connection 0 is used for this purpose.

### 27.7.15 CxpConnectionTestMode

|  Name | CxpConnectionTestMode[CxpConnectionSelector]  |
| --- | --- |
|  Category | CoaXPress  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off Model  |

Enables the test mode for an individual physical connection of the Device.

Possible values are: