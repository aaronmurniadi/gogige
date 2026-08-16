|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Size of the valid chunk payload data delivered in the buffer.

This information refers for example to the information provided in the GigE Vision image stream data trailer. For other technologies, this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_DELIVERED_CHUNKPAYLOADSIZE command of DSGetBufferInfo function.

### 3.5.2.28 BufferChunkLayoutID

|  Name | BufferChunkLayoutID  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

ID of the chunk data layout delivered in the buffer. Can be used to track changes of the layout data among individual buffers.

This information refers for example to the information provided in the GigE Vision image stream data leader. The chunk layout id serves as an indicator that the chunk layout has changed and the application should re-parse the chunk layout in the buffer. When a chunk layout (availability or position of individual chunks) changes since the last buffer delivered by the device through the same stream, the device MUST change the chunk layout id. As long as the chunk layout remains stable, the camera MUST keep the chunk layout id intact. When switching back to a layout, which was already used before, the camera can use the same id again or use a new id. A chunk layout id value of 0 is invalid. It is reserved for use by cameras not supporting the layout id functionality. The algorithm used to compute the chunk layout id is left as quality of implementation. For other technologies this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_CHUNKLAYOUTID command of DSGetBufferInfo function.

### 3.5.2.29 BufferFileName

|  Name | BufferFileName  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |