|  ![img-196.jpeg](img-196.jpeg)CAN |   | ![img-197.jpeg](img-197.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Acronym | Description  |
| --- | --- |
|  P/CM | The command is not available in case the buffer contains multi-part data. In this case the function DSGetBufferInfo returns GC_ERR_NO_DATA.In case the buffer does not contain multi-part data the command returns the requested information. In this case the implementation of the command is conditional mandatory. Conditional mandatory is used for commands which might not always be applicable. If it is possible to implement a certain command it must be implemented.  |

|  Enumerator | Global -Part /Impl | Value | Description  |
| --- | --- | --- | --- |
|  BUFFER_INFO_BASE | G/M | 0 | Base address of the buffer memory as passed to theDSAnnounceBuffer function or allocated byDSAllocAndAnnounceBuffer.This is also the address where the payload within the buffer starts.This info is irrelevant for composite buffers announced usingDSAnnounceCompositeBuffer, function returns GC_ERR_NO_DATA.Data type: PTR  |
|  BUFFER_INFO_SIZE | G/M | 1 | Size of the buffer in bytes.For composite buffers announced usingDSAnnounceCompositeBuffer, this returns the sum of sizes of the composite buffer segments.Data type: SIZET  |
|  BUFFER_INFO_USER_PTR | G/O | 2 | Private data pointer casted to an integer provided at buffer announcement using one of the buffer announcement functions by the GenTL Consumer.The pointer is attached to the buffer to allow attachment of user data to a buffer.Data type: PTR  |