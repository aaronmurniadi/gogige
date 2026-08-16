|  GEN<i>CAM |   | ![img-18.jpeg](img-18.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | DeviceEnumeration  |
| --- | --- |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Current subnet mask of the GVCP interface of the selected remote device. This value only changes on execution of the DeviceUpdateList command.

Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

3.2.2.15 GevDeviceGateway

|  Name | GevDeviceGateway[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Current gateway IP address of the GVCP interface of the selected remote device. This value only changes on execution of the DeviceUpdateList command.

3.2.2.16 GevDeviceIPConfigurationStatus

|  Name | GevDeviceIPConfigurationStatus[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Recommended  |