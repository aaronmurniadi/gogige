|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.3.1.16 GevDeviceMACAddress

|  Name | GevDeviceMACAddress  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

48-bit MAC address of the GVCP interface of the remote device.

Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory, and that the Representation element should be used in the XML to facilitate understanding the data.

### 3.3.1.17 GevDeviceIPAddress

|  Name | GevDeviceIPAddress  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Current IP address of the GVCP interface of the remote device.

Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.