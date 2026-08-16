|  ![img-198.jpeg](img-198.jpeg)CAN |   | ![img-199.jpeg](img-199.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Global -Part /Impl | Value | Description  |
| --- | --- | --- | --- |
|  BUFFER_INFO_TIMESTAMP | G/O | 3 | Timestamp the buffer was acquired. The unit is device/implementation dependent. In case the technology and/or the device does not support this for example under Windows a QueryPerformanceCounter can be used. Data type: UINT64  |
|  BUFFER_INFO_NEW_DATA | G/M | 4 | Flag to indicate that the buffer contains new data since the last delivery. Value 0 means this buffer has not been processed by the acquisition engine (it got delivered e.g. as a result of a flush operation). In such case other data related buffer info queries may fail. However, data agnostic information (such as BUFFER_INFO_BASE, BUFFER_INFO_SIZE, or BUFFER_INFO_USER_PTR) must still be correctly reported by GenTL Producer. Data type: BOOL8  |
|  BUFFER_INFO_IS_QUEUED | G/M | 5 | If this flag is set to true the buffer is in the input pool, the buffer is currently being filled or the buffer is in the output buffer queue. In case this value is true the buffer is owned by the GenTL Producer and it can not be revoked. Data type: BOOL8  |
|  BUFFER_INFO_IS_ACQUIRING | G/CM | 6 | Flag to indicate that the buffer is currently being filled with data. Data type: BOOL8  |
|  BUFFER_INFO_IS_INCOMPLETE | G/M | 7 | Flag to indicate that a buffer was filled but an error occurred during that process. Data type: BOOL8  |
|  BUFFER_INFO_TLTYPE | G/M | 8 | Transport layer technology that is supported. See string constants in chapter 6.6.1. Data type: STRING  |