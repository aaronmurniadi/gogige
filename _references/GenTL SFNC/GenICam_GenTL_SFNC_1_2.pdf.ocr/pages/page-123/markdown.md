|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

This information refers for example to the width entry in the GigE Vision image stream data leader. For other technologies, this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_WIDTH command of DSGetBufferInfo function.

### 3.5.2.17 BufferHeight

|  Name | BufferHeight[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Height of the data in the buffer in number of pixels as configured. For variable size images this is the max Height of the buffer.

For example this information refers to the height entry in the GigE Vision image stream data leader. For other technologies this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_HEIGHT command of DSGetBufferInfo function.

### 3.5.2.18 BufferXOffset

|  Name | BufferXOffset[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

XOffset of the data in the buffer in number of pixels from the image origin to handle areas of interest.