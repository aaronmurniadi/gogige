|  ![img-204.jpeg](img-204.jpeg) CAM |   | ![img-205.jpeg](img-205.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Global -Part /Impl | Value | Description  |
| --- | --- | --- | --- |
|  BUFFER_INFO_PIXELFORMAT_NAMESPACE | P/CM | 21 | This information refers to the constants defined in PIXELFORMAT NAMESPACE IDs to allow interpretation of BUFFER_INFO_PIXELFORMAT. Data type: UINT64  |
|  BUFFER_INFO_DELIVERED_IMAGEHEIGHT | P/CM | 22 | The number of lines in the current buffer as delivered by the transport mechanism. For area scan type images this is usually the number of lines configured in the device. For variable size linescan images this number may be lower than the configured image height. This information refers for example to the information provided in the GigE Vision image stream data trailer. For other technologies this is to be implemented accordingly. Data type: SIZET  |
|  BUFFER_INFO_DELIVERED_CHUNKPAYLOADSIZE | G/CM | 23 | This information refers for example to the information provided in the GigE Vision image stream data trailer. For other technologies this is to be implemented accordingly. For GenDC payload the eventual chunk data must be processed based on the information in the GenDC descriptor, the GenTL Producer must return GC_ERR_NO_DATA in that case. Data type: SIZET  |