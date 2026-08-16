|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.4.3.5 StreamDeliveredFrameCount

|  Name | StreamDeliveredFrameCount  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Number of delivered frames since last acquisition start. It is not reset until the stream is closed.

Corresponds to the STREAM_INFO_NUM_DELIVERED command of DSGetInfo function.

### 3.4.3.6 StreamLostFrameCount

|  Name | StreamLostFrameCount  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Number of lost frames due to queue underrun. This number is initialized with zero at the time the stream is opened and incremented every time the data could not be acquired because there was no buffer in the input buffer pool. It is not reset until the stream is closed.

Corresponds to the STREAM_INFO_NUM_UNDERRUN command of DSGetInfo function.

### 3.4.3.7 StreamInputBufferCount

|  Name | StreamInputBufferCount  |
| --- | --- |
|  Category | BufferHandlingControl  |