Sets the value of the bit selected by UserOutputSelector.

### 9.2.11 UserOutputValueAll

|  Name | UserOutputValueAll  |
| --- | --- |
|  Category | DigitalIOControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Device-specific  |

Sets the value of all the bits of the User Output register. It is subject to the UserOutputValueAllMask.

UserOutputValueAll can take any binary value and each bit set to one will set the corresponding User Output register bit to high. Note that the UserOutputs are numbered from 0 to N (If 0 based). This means that the least significant bit of UserOutputValueAll corresponds to the UserOutput0 (if 0 based).

### 9.2.12 UserOutputValueAllMask

|  Name | UserOutputValueAllMask  |
| --- | --- |
|  Category | DigitalIOControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Device-specific  |

Sets the write mask to apply to the value specified by UserOutputValueAll before writing it in the User Output register. If the UserOutputValueAllMask feature is present, setting the user Output register using UserOutputValueAll will only change the bits that have a corresponding bit in the mask set to one.

UserOutputValueAllMask can take any binary value. Each bit set to one will enable writing of the corresponding User Output register bit and each bit set to zero will prevent it.

Note that UserOutputValueAllMask is ignored when an individual bit is set using UserOutputValue.