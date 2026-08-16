|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 27.4.65 GevSCPSDoNotFragment

|  Name | GevSCPSDoNotFragment[GevStreamChannelSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | True False  |

The state of this feature is copied into the "do not fragment" bit of IP header of each stream packet. It can be used by the application to prevent IP fragmentation of packets on the stream channel.

### 27.4.66 GevSCPSBigEndian (Deprecated)

|  Name | GevSCPSBigEndian[GevStreamChannelSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | True False  |

This feature is deprecated (See DeviceStreamChannelEndianness). It was used to control the endianness of multi-byte pixel data for this stream.

This is an optional feature. A device that does not support this feature must support little-endian and always leave that bit clear.

### 27.4.67 GevSCPSPacketSize

|  Name | GevSCPSPacketSize[GevStreamChannelSelector]  |
| --- | --- |
|  Category | GigEVision  |