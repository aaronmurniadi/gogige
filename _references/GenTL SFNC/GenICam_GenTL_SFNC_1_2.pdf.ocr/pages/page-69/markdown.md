|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  TLType | GigEVision  |
| --- | --- |
|  Interface | IEnum  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | DHCP PersistentIP LinkLocal  |

Device IP configuration of the GVCP interface of the selected remote device. This value only changes on execution of the DeviceUpdateList command.

### 3.2.2.17 GevDeviceMACAddress

|  Name | GevDeviceMACAddress[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

48-bit MAC address of the GVCP interface of the selected remote device.

Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory, and that the Representation element should be used in the XML to facilitate understanding the data.

### 3.2.2.18 GevDeviceCurrentControlMode

|  Name | GevDeviceCurrentControlMode[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Optional  |