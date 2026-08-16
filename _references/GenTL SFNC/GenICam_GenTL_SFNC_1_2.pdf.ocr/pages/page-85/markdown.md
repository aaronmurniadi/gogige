|  **GEN<i>CAM** |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Unit | -  |
| --- | --- |
|  Visibility | Beginner  |
|  Values | >0  |

The tick-frequency of the time stamp clock.

Corresponds to the DEVICE_INFO_TIMESTAMP_FREQUENCY command of DevGetInfo function.

### 3.3.1.13 DeviceAccessStatus

|  Name | DeviceAccessStatus  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Unknown ReadWrite ReadOnly NoAccess Busy OpenReadWrite OpenReadOnly  |

Gives the device's access status at the moment of the last execution of the DeviceUpdateList command. This value only changes on execution of the DeviceUpdateList command.

- **Unknown**: Not known to producer.
- **ReadWrite**: Full access
- **ReadOnly**: Read-only access
- **NoAccess**: Not available to connect.
- **Busy**: The device is already opened by another entity.