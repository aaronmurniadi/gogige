|  ![img-206.jpeg](img-206.jpeg)CAN |   | ![img-207.jpeg](img-207.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Global -Part /Impl | Value | Description  |
| --- | --- | --- | --- |
|  BUFFER_INFO_FILENAME | G/CM | 25 | Filename in case the payload contains a file.This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies this is to be implemented accordingly. Since this is GigE Vision related information and the filename in GigE Vision is UTF8 coded, this filename is also UTF8 coded.Data type: STRING  |
|  BUFFER_INFO_PIXEL_ENDIANNESS | G/O | 26 | Endianness of the multi-byte pixel data in the buffer. This information refers to the constants defined in PIXELENDIANNESS_IDs.Data type: INT32  |
|  BUFFER_INFO_DATA_SIZE | G/O | 27 | Size of the data intended to be written to the buffer last time it has been filled. This value is reset to 0 when the buffer is placed into the Input Buffer Pool.If the buffer is incomplete the number still reports the full size of the original data including the lost parts. If the buffer is complete and the payload contains no gaps, the number equals to the size reported through BUFFER_INFO_SIZE_FILLED.Note that gaps in payload are possible with certain payload types, such as GenDC or multi-part.Data type: SIZET  |
|  BUFFER_INFO_TIMESTAMP_NS | G/O | 28 | Timestamp the buffer was acquired, in units of 1 ns (1 000 000 000 ticks per second). If the device is internally using another tick frequency than 1GHz, the GenTL Producer must convert the value to nanoseconds.Data type: UINT64  |