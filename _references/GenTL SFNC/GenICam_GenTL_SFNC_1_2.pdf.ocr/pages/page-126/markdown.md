|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

This information refers for example to the information provided in the GigE Vision image stream block id. For other technologies, this is to be implemented accordingly. The wrap around of this number is transportation technology dependent

Corresponds to the BUFFER_INFO_FRAMEID command of DSGetBufferInfo function.

### 3.5.2.23 BufferImagePresent

|  Name | BufferImagePresent  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Flag to indicate if the current data in the buffer contains image data.

This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies, this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_IMAGEPRESET command of DSGetBufferInfo function.

### 3.5.2.24 BufferImageOffset

|  Name | BufferImageOffset  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Byte  |
|  Visibility | Expert  |
|  Values | ≥0  |