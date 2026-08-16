|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies, this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_XPADDING command of DSGetBufferInfo function.

### 3.5.2.21 BufferYPadding

|  Name | BufferYPadding  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Byte  |
|  Visibility | Expert  |
|  Values | ≥0  |

YPadding of the data in the buffer in number of bytes.

This information refers for example to the information provided in the GigE Vision image stream data leader. For other thechnologies, this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_YPADDING command of DSGetBufferInfo function.

### 3.5.2.22 BufferFrameID

|  Name | BufferFrameID  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

A sequentially incremented number of the frame.