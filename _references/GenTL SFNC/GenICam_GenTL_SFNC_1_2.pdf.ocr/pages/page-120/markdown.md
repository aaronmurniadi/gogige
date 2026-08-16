|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Level | Optional  |
| --- | --- |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

The buffer part to extract information from. For non-multipart the value is 0. The maximum value should be dynamic and reflect the number of parts possible to index.

##### 3.5.2.11 BufferSizeFilled

|  Name | BufferSizeFilled  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Byte  |
|  Visibility | Expert  |
|  Values | ≥0  |

Number of bytes written into the buffer last time it was filled. This value is reset to 0 when the buffer is placed into the Input Buffer Pool.

Corresponds to the BUFFER_INFO_SIZE_FILLED command of DSGetBufferInfo function.

##### 3.5.2.12 BufferPartDataType

|  Name | BufferPartDataType[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |