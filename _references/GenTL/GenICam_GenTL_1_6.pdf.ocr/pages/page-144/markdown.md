|  ![img-202.jpeg](img-202.jpeg) CAM |   | ![img-203.jpeg](img-203.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Global -Part /Impl | Value | Description  |
| --- | --- | --- | --- |
|  BUFFER_INFO_IMAGEPRESENT | G/M | 17 | Flag to indicate if the current data in the buffer contains single raster scan image data. This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies this is to be implemented accordingly.The flag should report true if the payload contains a single raster scan image in the sense of traditional image payload type and if it makes sense to query its properties using buffer-global image format related buffer infos.Data type: BOOL8  |
|  BUFFER_INFO_IMAGEOFFSET | G/O | 18 | Offset of the image data from the beginning of the delivered buffer in bytes. Applies for example when delivering the image as part of chunk data or on technologies requiring specific buffer alignment.Data type: SIZET  |
|  BUFFER_INFO_PAYLOADTYPE | G/M | 19 | Payload type of the data.This information refers to the constants defined inPAYLOADTYPE_INFO_IDS.Data type: SIZET  |
|  BUFFER_INFO_PIXELFORMAT | P/CM | 20 | Pixelformat of the data.This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies this is to be implemented accordingly. The interpretation of the pixel format depends on the namespace the pixel format belongs to. This can be inquired using theBUFFER_INFO_PIXELFORMAT_NAMESPACE command.Data type: UINT64  |