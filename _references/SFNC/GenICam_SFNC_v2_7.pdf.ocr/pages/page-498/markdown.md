|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

## 25 Test Control

Contains the features related to the control of the test features.

### 25.1 TestControl

|  Name | TestControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | -  |

Category for Test Control features.

### 25.2 TestPendingAck

|  Name | TestPendingAck  |
| --- | --- |
|  Category | TestControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | ms  |
|  Visibility | Guru  |
|  Values | ≥0  |

Tests the device's pending acknowledge feature. When this feature is written, the device waits a time period corresponding to the value of TestPendingAck before acknowledging the write.

If this time period is longer than the maximum device response time specified by DeviceLinkCommandTimeout, the device must use a pending acknowledge during the completion of this request.

When read, the device returns the current feature value without additional wait time before the acknowledge.