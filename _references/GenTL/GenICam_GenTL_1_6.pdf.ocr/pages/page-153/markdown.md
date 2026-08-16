|  ![img-217.jpeg](img-217.jpeg) CAM |   | ![img-218.jpeg](img-218.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|  STREAM_INFO_IS_GRABBING | M | 8 | Flag indicating whether the acquisition engine is started or not. This is independent from the acquisition status of the remote device.Data type: BOOL8  |
|  STREAM_INFO_DEFINES_PAYLOADSIZE | M | 9 | Flag indicating that this data stream defines a payload size independent from the remote device. If false the size of the expected payload size is to be retrieved from the remote device. If true the expected payload size is to be inquired from the Data Stream module. In case the GenTL Producer does not support this command it is to interpreted as false.Data type: BOOL8  |
|  STREAM_INFO_TLTYPE | M | 10 | Transport layer technology that is supported. See string constants in chapter 6.6.1.Data type: STRING  |
|  STREAM_INFO_NUM_CHUNKS_MAX | CM | 11 | Maximum number of chunks to be expected in a buffer (can be used to allocate the array for the DSGetBufferChunkData function). In case this is not known a priori by the GenTL Producer the DSGetInfo function returns GC_ERR_NOT_AVAILABLE. This maximum must not change during runtime.Data type: SIZET  |
|  STREAM_INFO_BUF_ANNOUNCE_MIN | M | 12 | Minimum number of buffers to announce. In case this is not known a priori by the GenTL Producer the DSGetInfo function returns a GC_ERR_NOT_AVAILABLE error. This minimum may change during runtime when changing parameters through the node map.Data type: SIZET  |
|  STREAM_INFO_BUF_ALIGNMENT | O | 13 | Alignment size in bytes of the buffer base pointer passed to DSAnnounceBuffer.If a buffer is passed to  |