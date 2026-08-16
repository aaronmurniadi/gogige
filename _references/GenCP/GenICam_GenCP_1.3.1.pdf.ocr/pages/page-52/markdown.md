|  GEN<i>CAM |   | ![img-55.jpeg](img-55.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 5.4.7. Serial Number

The register contains a string representing the serial number of the device.

|  Offset | Hex 144  |
| --- | --- |
|  Length | 64  |
|  Access Type | R  |
|  Support | M  |
|  Data Type | String  |
|  Factory Default | Device specific  |

#### 5.4.8. User Defined Name

A string containing a user defined name. A write to this register must instantly persist without explicitly being stored to non-volatile memory. The User Defined Name Supported bit in the Device Capability register indicates if this register is present or not.

|  Offset | Hex 184  |
| --- | --- |
|  Length | 64  |
|  Access Type | RW  |
|  Support | CM  |
|  Data Type | String  |
|  Factory Default | Empty String  |