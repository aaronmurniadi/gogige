|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_XOFFSET command of DSGetBufferInfo function.

### 3.5.2.19 BufferYOffset

|  Name | BufferYOffset[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

YOffset of the data in the buffer in number of lines from the image origin to handle areas of interest.

This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies, this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_YOFFSET command of DSGetBufferInfo function.

### 3.5.2.20 BufferXPadding

|  Name | BufferXPadding[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Byte  |
|  Visibility | Expert  |
|  Values | ≥0  |

XPadding of the data in the buffer in number of bytes.