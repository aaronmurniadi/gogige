|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.4.3 Buffer Handling Control

Features in this section provide control over the buffers within the acquisition engine of a data stream.

#### 3.4.3.1 BufferHandlingControl

|  Name | BufferHandlingControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Contains all features of the Data Stream module that control the used buffers.

#### 3.4.3.2 StreamAnnouncedBufferCount

|  Name | StreamAnnouncedBufferCount  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Number of announced (known) buffers on this stream. This value is volatile. It may change if additional buffers are announced and/or buffers are revoked by the GenTL Consumer.

Corresponds to the STREAM_INFO_NUM_ANNOUNCED command of DSGetInfo function.

#### 3.4.3.3 StreamBufferHandlingMode

|  Name | StreamBufferHandlingMode  |
| --- | --- |