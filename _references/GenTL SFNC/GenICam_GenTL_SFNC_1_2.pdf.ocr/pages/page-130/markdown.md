|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Convention (PFNC)" specification hosted by the AIA organisation. Refer to the most recent version of that convention for additional information about the construction of a pixel format name.

### 3.5.2.26 BufferDeliveredImageHeight

|  Name | BufferDeliveredImageHeight[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

The number of lines in the current buffer part as delivered by the transport mechanism. For area scan type images, this is usually the number of lines configured in the device. For variable size linescan images, this number may be lower than the configured image height.

This information refers for example to the information provided in the GigE Vision image stream data trailer. For other technologies, this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_DELIVERED_IMAGEHEIGHT command of DSGetBufferInfo function and BUFFER_PART_INFO_DELIVEREDIMAGEHEIGHT in a DSGetPartInfo function

### 3.5.2.27 BufferDeliveredChunkPayloadSize

|  Name | BufferDeliveredChunkPayloadSize  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |