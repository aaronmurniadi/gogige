|  ![img-223.jpeg](img-223.jpeg)CAN |   | ![img-224.jpeg](img-224.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|   |  |  | depends on the namespace the format belongs to which can be inquired using theBUFFER_PART_INFO_DATA_FORMAT_NAMESPACE command (although for the standardPARTDATATYPE_IDS arecommended data format namespace is always specified).Data type: UINT64  |
|  BUFFER_PART_INFO_DATA_FORMAT_NAMESPACE | M | 4 | This information refers to the constants defined inPIXELFORMAT_NAMESPACE_IDSto allow interpretation ofBUFFER_PART_INFO_DATA_FORMAT.Data type: UINT64  |
|  BUFFER_PART_INFO_WIDTH | CM | 5 | Width of the data in the buffer part in number of pixels.If the information is not applicable to given data type, the query should result in GC_ERR_NOT_AVAILABLE.Data type: SIZET  |
|  BUFFER_PART_INFO_HEIGHT | CM | 6 | Height of the data in the buffer part in number of pixels.If the information is not applicable to given data type, the query should result in GC_ERR_NOT_AVAILABLE.Data type: SIZET  |
|  BUFFER_PART_INFO_XOFFSET | CM | 7 | XOffset of the data in the buffer part in number of pixels from the image origin to handle areas of interest.If the information is not applicable to given data type, the query should result in GC_ERR_NOT_AVAILABLE.Data type: SIZET  |
|  BUFFER_PART_INFO_YOFFSET | CM | 8 | YOffset of the data in the buffer part in number of pixels from the image origin  |