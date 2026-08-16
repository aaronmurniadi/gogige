|  ![img-212.jpeg](img-212.jpeg) CAM |   | ![img-213.jpeg](img-213.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  PAYLOAD_TYPE_JPEG2000 | 6 | The buffer payload contains JPEG 2000 data in the format described in GEV 2.0. The GenTL Producer should report additional information through the corresponding BUFFER_INFO_CMD commands.  |
|  PAYLOAD_TYPE_H264 | 7 | The buffer payload contains H.264 data in the format described in GEV 2.0. The GenTL Producer should report additional information through the corresponding BUFFER_INFO_CMD commands.  |
|  PAYLOAD_TYPE_CHUNK_ONLY | 8 | The buffer payload contains only chunk data but no additional payload.  |
|  PAYLOAD_TYPE_DEVICE_SPECIFIC | 9 | The buffer payload contains device specific data. The GenTL Producer should report additional information through the corresponding BUFFER_INFO_CMD commands.  |
|  PAYLOAD_TYPE_MULTI_PART | 10 | The buffer payload contains multiple parts of different payload types.Information about the individual parts should be queried using DSGetNumBufferParts and DSGetBufferPartInfo functions.  |
|  PAYLOAD_TYPE_GENDC | 11 | The buffer payload contains a GenDC container. Its contents must be interpreted based on the rules defined in GenDC specification.GenTL specific details related to GenDC payload transfer are defined in 5.7.3.  |
|  PAYLOAD_TYPE_CUSTOM_ID | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific.  |

##### 6.4.4.6 PIXELFORMAT_NAMESPACE_IDS

enum PIXELFORMAT_NAMESPACE_IDS

This enumeration defines constants to interpret the pixel formats provided through BUFFER_INFO_PIXELFORMAT.

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  PIXELFORMAT_NAMESPACE_UNKNOWN | 0 | The interpretation of the pixel format values is unknown to the GenTL Producer.  |
|  PIXELFORMAT_NAMESPACE_GEV | 1 | The interpretation of the pixel  |