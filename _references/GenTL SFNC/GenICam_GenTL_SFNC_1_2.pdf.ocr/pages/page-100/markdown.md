|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.4.2.5 DeviceStreamChannelPacketSizeInc

|  Name | DeviceStreamChannelPacketSizeInc  |
| --- | --- |
|  Category | DeviceStreamChannelControl  |
|  Level | Optional  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | B  |
|  Visibility | Guru  |
|  Values | >0  |

Controls desired increment of the packet size feature to be used for the stream channel. Affects both the direct control of the packet size as well as the negotiation algorithm. The GenTL Consumer can set the value in accordance with the known limits of the remote device or apply further restrictions e.g. based on additional knowledge of the system.

### 3.4.2.6 DeviceStreamChannelNegotiatePacketSize

|  Name | DeviceStreamChannelNegotiatePacketSize  |
| --- | --- |
|  Category | DeviceStreamChannelControl  |
|  Level | Optional  |
|  TLType | GigEVision  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Starts negotiation for the optimal packet size considering the remote device, host and their connection path. The negotiation result is applied on the device and reflected in DeviceStreamChannelPacketSize. If the negotiation fails, the algorithm attempts to revert the configuration to the initial packet size value.