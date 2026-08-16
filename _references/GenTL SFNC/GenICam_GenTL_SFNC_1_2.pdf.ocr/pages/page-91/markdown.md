|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Default value is “Legacy”.

Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory, otherwise recommended.

### 3.3.2.3 LinkCommandTimeout

|  Name | LinkCommandTimeout  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Recommended  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | us  |
|  Visibility | Guru  |
|  Values | >0  |

Specifies application timeout for the control channel communication. This feature defines the application timeout, and it is related to the device feature DeviceLinkCommandTimeout specifying the maximum time for handling a command in the device. Up to DeviceLinkCommandRetryCount attempts with this timeout are made before a command fails with a timeout error.

### 3.3.2.4 LinkCommandRetryCount

|  Name | LinkCommandRetryCount  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | >=0  |

Specifies maximum number of tries before failing the control channel commands.