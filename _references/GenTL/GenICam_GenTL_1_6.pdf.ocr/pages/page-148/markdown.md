|  ![img-208.jpeg](img-208.jpeg)CAM |   | ![img-209.jpeg](img-209.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Global -Part /Impl | Value | Description  |
| --- | --- | --- | --- |
|  BUFFER_INFO_DATA_LARGER_THAN_BUFFER | G/O | 29 | If this values is set to true it indicates that the payload transferred would not fit into the announced buffer and that therefore only parts of the payload or no payload (depending on the implementation of the GenTL Producer) is copied into the buffer. Data type: BOOL8  |
|  BUFFER_INFO_CONTAINS_CHUNKDATA | G/M | 30 | If this values is set to true it indicates that the payload transferred contains chunk data which may be parsed through a call to DSGetBufferChunkData or the GenTL Consumer.For GenDC payload the eventual chunk data must be processed based on the information in the GenDC descriptor, the GenTL Producer must return GC_ERR_NO_DATA in that case. Data type: BOOL8  |
|  BUFFER_INFO_IS_COMPOSITE | G/M | 31 | Indicates whether this is a composite buffer, announced using DSAnnounceCompositeBuffer. Data type: BOOL8  |
|  BUFFER_INFO_CUSTOM_ID | G/O | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific.If a generic GenTL Consumer is using custom BUFFER_INFO_CMDs provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |