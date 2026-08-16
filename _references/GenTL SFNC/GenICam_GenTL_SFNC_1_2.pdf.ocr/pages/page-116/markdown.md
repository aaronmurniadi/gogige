|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | BufferInformation  |
| --- | --- |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Timestamp the buffer was acquired. The unit is device/implementation dependent. In case the technology and/or the device does not support this for example under Windows a QueryPerformanceCounter can be used.

Corresponds to the BUF_INFO_TIMESTAMP command of DSGetBufferInfo function.

### 3.5.2.4 BufferNewData

|  Name | BufferNewData  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Flag to indicate that the buffer contains new data since the last delivery.

Corresponds to the BUFFER_INFO_NEW_DATA command of DSGetBufferInfo function.

### 3.5.2.5 BufferIsQueued

|  Name | BufferIsQueued  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IBoolean  |