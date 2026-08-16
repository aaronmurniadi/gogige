|  ![img-239.jpeg](img-239.jpeg) CAM |   | ![img-240.jpeg](img-240.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|   |  |  | If a generic GenTL Consumer is using custom FLOW_INFO_CMDs provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |

##### 6.4.4.12 SEGMENT_INFO_CMD

enum SEGMENT_INFO_CMD

This enumeration defines commands to retrieve information with the DSGetBufferSegmentInfo function on a buffer handle.

The column labeled “Impl” in the following table lists if the implementation of a given command is mandatory (M), optional (O) or conditional mandatory (CM). Mandatory means that a GenTL Producer must implement the listed command. Optional means that it is up to the implementor if a given command is implemented or not. Conditional Mandatory means that command is to be implemented if possible.

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|  SEGMENT_INFO_BASE | M | 0 | Base address of the buffer segment memory as passed to theDSAnnounceCompositeBuffer function. This is also the address where the payload within the segment starts.Data type: PTR  |
|  SEGMENT_INFO_SIZE | M | 1 | Size of the buffer segment in bytes as passed to theDSAnnounceCompositeBuffer function.Data type: SIZET  |
|  SEGMENT_INFO_IS_INCOMPLETE | O | 2 | Flag to indicate that the buffer segment was filled but an error occurred during that process.For technologies or use cases where this is difficult to track, it is valid leave the command not implemented.Data type: BOOL8  |
|  SEGMENT_INFO_SIZE_FILLED | O | 3 | Number of bytes written into the buffer segment the last time it has been filled.This value is reset to 0 when the buffer  |