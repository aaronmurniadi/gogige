Version 1.6

GenTL Standard

|  Enumerator | Global -Part /Impl | Value | Description  |
| --- | --- | --- | --- |
|  BUFFER_INFO_YOFFSET | P/CM | 13 | YOffset of the data in the buffer in number of lines from the image origin to handle areas of interest. This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies this is to be implemented accordingly.Data type: SIZET  |
|  BUFFER_INFO_XPADDING | P/CM | 14 | XPadding of the data in the buffer in number of bytes. This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies this is may be implemented accordingly.Data type: SIZET  |
|  BUFFER_INFO_YPADDING | G/O | 15 | YPadding of the data in the buffer in number of bytes. This information refers for example to the information provided in the GigE Vision image stream data leader. For other thechnologies this may be implemented accordingly.Data type: SIZET  |
|  BUFFER_INFO_FRAMEID | G/M | 16 | A sequentially incremented number of the frame. This information refers for example to the information provided in the GigE Vision image stream block id. For other technologies this is to be implemented accordingly.The wrap around of this number is transportation technology dependent. For GigE Vision it is (so far) 16bit wrapping to 1. Other technologies may implement a larger bit depth.Data type: UINT64  |