|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 2.5 Buffer Module

#### 2.5.1 Buffer Information

Contains the features related to general information about a specific Buffer module.

Table 2-20: Buffer Information Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  BufferInformation | O | All | ICategory | R | - | E | Category that contains all Buffer Information features of the Buffer module.  |
|  BufferUserData | O | All | IInteger | R | - | E | Pointer to user data casted to an integer number referencing GenTL Consumer specific data.  |
|  BufferType | O | All | IEnumeration | R | - | E | Transport layer type of the buffer.  |
|  BufferSize | O | All | IInteger | R | Byte | E | Size of the buffer in bytes.  |

#### 2.5.2 Buffer Data Information

Contains the features related to the currently filled data of a specific Buffer module.

Table 2-21: Buffer Data Information Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  BufferDataInformation | O | All | ICategory | R | - | E | Contains all Buffer Data Information features of the Buffer module.  |