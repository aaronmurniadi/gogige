|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | None  |
| --- | --- |
|  Level | Optional  |
|  Interface | IPort  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | -  |

The GenICam port through which the Buffer module is accessed.

Note that BufferPort is a port node (not a feature node) and is generally not accessed by the end user directly.

Note that according to the GenICam GenTL standard, this feature is not mandatory. However, if this feature is provided, also the features “BufferData” and “BufferUserData” are mandatory.