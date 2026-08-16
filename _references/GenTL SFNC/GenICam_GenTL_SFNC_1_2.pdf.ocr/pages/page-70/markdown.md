|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  TLType | GigEVision  |
| --- | --- |
|  Interface | IEnum  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Open ControlAccess ExclusiveAccess NoAccess  |

The current control mode of the device. This value only changes on execution of the DeviceUpdateList command. See also DeviceAccessStatus, which gives a similar TL independent status. The values are.

- Open : The device is open for control or exclusive access.
- ControlAccess: The device is controlled by another host, but switchover or readonly access is possible.
- ExclusiveAccess: The device is under exclusive access by a host and cannot be accessed by another.
- NoAccess: The device cannot be accessed, for instance it may be a GigE Vision device on a subnet different from the interface.

### 3.2.2.19 GevApplicationSwitchoverKey

|  Name | GevApplicationSwitchoverKey[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Optional  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |