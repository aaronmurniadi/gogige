|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

## 2.4 Data Stream Module

### 2.4.1 Stream Information

Contains the features related to general information about a specific Data Stream module.

Table 2-15: Stream Information Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  StreamInformation | M | All | ICategory | R | - | E | Category that contains all Stream Information features of the Data Stream module.  |
|  StreamID | M | All | IString | R | - | E | Device unique ID for the data stream.  |
|  StreamType | M | All | IEnumeration | R | - | E | Transport layer type of the Data Stream.  |

### 2.4.2 Device Stream Channel Control

Contains the features related to control the buffers within the acquisition engine of a specific Data Stream module.

Table 2-16: Buffer Handling Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  DeviceStreamChannelControl | R | GEV | ICategory | R | - | E | Category containing features to control the stream channel shared between the remote device and the GenTL Producer's data stream module.  |
|  DeviceStreamChannelPacketSize | R | GEV | IInteger | R/(W) | B | E | Specifies the stream packet size, in bytes, to send on the selected channel for a transmitter or specifies the maximum packet size supported by a receiver.  |
|  DeviceStreamChannelPacketSizeMin | O | GEV | IInteger | R/(W) | B | G | Controls desired minimum of the packet size feature to be  |