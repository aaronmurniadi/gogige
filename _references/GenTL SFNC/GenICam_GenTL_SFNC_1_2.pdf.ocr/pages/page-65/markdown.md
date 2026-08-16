|  GEN<i>CAM |   | ![img-17.jpeg](img-17.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Interface | IEnumeration  |
| --- | --- |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | UnknownReadWriteReadOnlyNoAccessBusyOpenReadWriteOpenReadOnly  |

Gives the device's access status at the moment of the last execution of the DeviceUpdateList command. This value only changes on execution of the DeviceUpdateList command.

- Unknown : Not known to producer.
- ReadWrite: Full access
- ReadOnly: Read-only access
- NoAccess: Not available to connect.
- Busy: The device is already opened by another entity.
- OpenReadWrite : Open in Read/Write mode by this GenTL host
- OpenReadOnly : Open in Read only mode by this GenTL host

3.2.2.9 DeviceSerialNumber

|  Name | DeviceSerialNumber[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |