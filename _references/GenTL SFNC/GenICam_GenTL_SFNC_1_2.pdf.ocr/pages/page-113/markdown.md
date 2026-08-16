|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Note that this category depends whether a Port access is provided through the "BufferPort" feature.

### 3.5.1.2 BufferUserData

|  Name | BufferUserData  |
| --- | --- |
|  Category | BufferInformation  |
|  Level | Optional (but mandatory if Port access provided)  |
|  Interface | Integer  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Pointer to user data casted to an integer number referencing GenTL Consumer specific data. It is reflecting the pointer provided by the user data pointer at buffer announcement. This allows the GenTL Consumer to attach information to a buffer.

Note that according to the GenICam GenTL standard, this feature is mandatory if a Port access is provided through the “BufferPort” feature.

Corresponds to the BUFFER_INFO_USER_PTR command of DSGetBufferInfo function.

### 3.5.1.3 BufferType

|  Name | BufferType  |
| --- | --- |
|  Category | BufferInformation  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | GigEVision CameraLink CameraLinkHS CoaXPress USB3Vision  |