|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.4.2.3 DeviceStreamChannelPacketSizeMin

|  Name | DeviceStreamChannelPacketSizeMin  |
| --- | --- |
|  Category | DeviceStreamChannelControl  |
|  Level | Optional  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | B  |
|  Visibility | Guru  |
|  Values | >0  |

Controls desired minimum of the packet size feature to be used for the stream channel. Affects both the direct control of the packet size as well as the negotiation algorithm. The GenTL Consumer can set the value in accordance with the known limits of the remote device or apply further restrictions e.g. based on additional knowledge of the system.

### 3.4.2.4 DeviceStreamChannelPacketSizeMax

|  Name | DeviceStreamChannelPacketSizeMax  |
| --- | --- |
|  Category | DeviceStreamChannelControl  |
|  Level | Optional  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | B  |
|  Visibility | Guru  |
|  Values | >0  |

Controls desired maximum of the packet size feature to be used for the stream channel. Affects both the direct control of the packet size as well as the negotiation algorithm. The GenTL Consumer can set the value in accordance with the known limits of the remote device or apply further restrictions e.g. based on additional knowledge of the system.