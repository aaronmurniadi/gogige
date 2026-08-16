|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Level | Optional  |
| --- | --- |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Number of buffers in the input buffer pool plus the buffers(s) currently being filled.

Corresponds to the STREAM_INFO_NUM_QUEUED command of DSGetInfo function.

##### 3.4.3.8 StreamOutputBufferCount

|  Name | StreamOutputBufferCount  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Number of buffers in the output buffer queue.

Corresponds to the STREAM_INFO_NUM_AWAIT_DELIVERY command of DSGetInfo function.

##### 3.4.3.9 StreamStartedFrameCount

|  Name | StreamStartedFrameCount  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |