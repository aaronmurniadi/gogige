|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.4.2 Device Stream Channel Control

3.4.2.1 DeviceStreamChannelControl

|  Name | DeviceStreamChannelControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category containing features to control the stream channel shared between the remote device and the GenTL Producer's data stream module. Applicable for GigE Vision stream channels, and operating on the boot strap registers of the device since the nodemap for the device is not accessible to the GenTL producer.

3.4.2.2 DeviceStreamChannelPacketSize

|  Name | DeviceStreamChannelPacketSize  |
| --- | --- |
|  Category | DeviceStreamChannelControl  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | B  |
|  Visibility | Expert  |
|  Values | >0  |

Specifies the stream packet size, in bytes, to send on the selected channel for a transmitter or specifies the maximum packet size supported by a receiver. Controls the packet size configuration of the remote device and if needed the GenTL Producer.