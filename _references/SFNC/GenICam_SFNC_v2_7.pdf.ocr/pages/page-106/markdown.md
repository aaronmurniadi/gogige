|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|   | Little  |
| --- | --- |

Endianness of multi-byte pixel data for this stream.

Possible values are:

- Big: Stream channel data is big Endian.
- Little: Stream channel data is little Endian.

3.47 DeviceStreamChannelPacketSize

|  Name | DeviceStreamChannelPacketSize[DeviceStreamChannelSelector]  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | B  |
|  Visibility | Expert  |
|  Values | >0  |

Specifies the stream packet size, in bytes, to send on the selected channel for a Transmitter or specifies the maximum packet size supported by a receiver.

3.48 DeviceEventChannelCount

|  Name | DeviceEventChannelCount  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Indicates the number of event channels supported by the device.

3.49 DeviceMessageChannelCount (Deprecated)

|  Name | DeviceMessageChannelCount  |
| --- | --- |
|  Category | DeviceControl  |