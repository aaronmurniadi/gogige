|  ![img-200.jpeg](img-200.jpeg)CAN |   | ![img-201.jpeg](img-201.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Global -Part /Impl | Value | Description  |
| --- | --- | --- | --- |
|  BUFFER_INFO_SIZE_FILLED | G/O | 9 | Number of bytes written into the buffer the last time it has been filled. This value is reset to 0 when the buffer is placed into the Input Buffer Pool. If the buffer is incomplete (such as if there are missing packets), only the number of bytes successfully written to the buffer are reported. If the buffer is complete and payload contains no gaps, the number equals to the size reported through BUFFER_INFO_DATA_SIZE. Note that gaps in payload are possible with certain payload types, such as GenDC or multi-part. Data type: SIZET  |
|  BUFFER_INFO_WIDTH | P/CM | 10 | Width of the data in the buffer in number of pixels. This information refers for example to the width entry in the GigE Vision image stream data leader. For other technologies this is to be implemented accordingly. Data type: SIZET  |
|  BUFFER_INFO_HEIGHT | P/CM | 11 | Height of the data in the buffer in number of pixels as configured. For variable size images this is the maximum height of the buffer. For example this information refers to the height entry in the GigE Vision image stream data leader. For other technologies this is to be implemented accordingly. Data type: SIZET  |
|  BUFFER_INFO_XOFFSET | P/CM | 12 | XOffset of the data in the buffer in number of pixels from the image origin to handle areas of interest. This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies this is to be implemented accordingly. Data type: SIZET  |