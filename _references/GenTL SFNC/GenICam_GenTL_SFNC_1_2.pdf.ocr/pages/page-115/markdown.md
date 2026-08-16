|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.5.2.1 BufferDataInformation

|  Name | BufferDataInformation  |
| --- | --- |
|  Category | Root  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Contains all Buffer Data Information features of the Buffer module.

Note that this category depends whether a Port access is provided through the "BufferPort" feature.

### 3.5.2.2 BufferData

|  Name | BufferData  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional (but mandatory if Port access provided)  |
|  Interface | IRegister  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Entire buffer data.

Note that according to the GenICam GenTL standard, this feature is mandatory if a Port access is provided through the "BufferPort" feature.

Corresponds to the BUFFER_INFO_BASE command of DSGetBufferInfo function.

### 3.5.2.3 BufferTimeStamp

|  Name | BufferTimeStamp  |
| --- | --- |