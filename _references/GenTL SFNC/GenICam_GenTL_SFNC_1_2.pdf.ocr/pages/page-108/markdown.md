|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Values | True False  |
| --- | --- |

Flag indicating whether the acquisition engine is started or not. This is independent from the acquisition status of the remote device.

Corresponds to the STREAM_INFO_IS_GRABBING command of DSGetInfo function.

### 3.4.3.12 StreamChunkCountMaximum

|  Name | StreamChunkCountMaximum  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit |   |
|  Visibility | Expert  |
|  Values | ≥0  |

Maximum number of chunks to be expected in a buffer (can be used to allocate the array for the DSGetBufferChunkData function).

Corresponds to the STREAM_INFO_NUM_CHUNKS_MAX command of DSGetInfo function.

### 3.4.3.13 StreamBufferAlignment

|  Name | StreamBufferAlignment  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Byte  |
|  Visibility | Expert  |
|  Values | ≥0  |

Alignment size in bytes of the buffers passed to DSAnnounceBuffer.