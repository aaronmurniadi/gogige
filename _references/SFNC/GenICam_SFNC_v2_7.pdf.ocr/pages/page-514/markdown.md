|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 27.3.3 PtpClockAccuracy

|  Name | PtpClockAccuracy  |
| --- | --- |
|  Category | PtpControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Within25ns Within100ns Within250ns Within1us Within2p5us Within10us Within25us Within100us Within250us Within1ms Within2p5ms Within10ms Within25ms Within100ms Within250ms Within1s Within10s GreaterThan10s AlternatePTPProfile Unknown Reserved  |

Indicates the expected accuracy of the device PTP clock when it is the grandmaster, or in the event it becomes the grandmaster.

### 27.3.4 PtpDataSetLatch

|  Name | PtpDataSetLatch  |
| --- | --- |
|  Category | PtpControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |