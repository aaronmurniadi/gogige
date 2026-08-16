|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

It is necessary to latch the 64-bit timestamp value to guarantee its integrity when performing the two 32-bit read accesses to retrieve the higher and lower 32-bit portions.

### 27.4.37 GevDiscoveryAckDelay

|  Name | GevDiscoveryAckDelay  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | ms  |
|  Visibility | Expert  |
|  Values | ≥0 and ≤1000  |

Indicates the maximum randomized delay the device will wait to acknowledge a discovery command.

### 27.4.38 GevIEEE1588 (Deprecated)

|  Name | GevIEEE1588  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | True False  |

This feature is deprecated (See PtpEnable). It was used to enable the IEEE 1588 Precision Time Protocol to control the timestamp register.

### 27.4.39 GevIEEE1588ClockAccuracy (Deprecated)

|  Name | GevIEEE1588ClockAccuracy  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |