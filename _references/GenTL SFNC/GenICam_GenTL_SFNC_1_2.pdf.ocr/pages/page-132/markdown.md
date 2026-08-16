|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Interface | IString  |
| --- | --- |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Filename for the file payload data delivered in the buffer.

This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies, this is to be implemented accordingly. Since this is GigE Vision related information and the filename in GigE Vision is UTF8 coded, this filename is also UTF8 coded.

Corresponds to the BUFFER_INFO_FILENAME command of DSGetBufferInfo function.

#### 3.5.3 GenICam Control

This chapter provides the necessary features to use the GenICam feature tree of the Buffer module.

Note: In case of discrepancy between the features described in this chapter and the “GenICam Standard text” the GenTL SFNC document prevails.

3.5.3.1 Root

|  Name | Root  |
| --- | --- |
|  Category | None  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Provides the Root of the GenICam features tree.

3.5.3.2 BufferPort

|  Name | BufferPort  |
| --- | --- |