|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 27.4.32 GevHeartbeatTimeout (Deprecated)

|  Name | GevHeartbeatTimeout  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | ms  |
|  Visibility | Guru  |
|  Values | >0  |

This feature is deprecated (See DeviceLinkHeartbeatTimeout). It was controlling the current heartbeat timeout in milliseconds.

### 27.4.33 GevTimestampTickFrequency (Deprecated)

|  Name | GevTimestampTickFrequency  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Hz  |
|  Visibility | Invisible  |
|  Values | ≥0  |

This feature is deprecated (See the increment of the TimestampLatchValue feature). It was used to indicate the number of timestamp ticks in 1 second (frequency in Hz). If PTP is used, this feature must return 1,000,000,000 (1 GHz).

This is a 64 bits number.

### 27.4.34 GevTimestampControlLatch (Deprecated)

|  Name | GevTimestampControlLatch  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | ICommand  |