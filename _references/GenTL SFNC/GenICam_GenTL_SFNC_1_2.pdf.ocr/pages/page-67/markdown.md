|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Major version number of the transport layer specification the remote device complies with.

### 3.2.2.12 DeviceTLVersionMinor

|  Name | DeviceTLVersionMinor[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Minor version number of the transport layer specification the remote device complies with.

### 3.2.2.13 GevDeviceIPAddress

|  Name | GevDeviceIPAddress[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Current IP address of the GVCP interface of the selected remote device. This value only changes on execution of the DeviceUpdateList command.

Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

### 3.2.2.14 GevDeviceSubnetMask

|  Name | GevDeviceSubnetMask[DeviceSelector]  |
| --- | --- |