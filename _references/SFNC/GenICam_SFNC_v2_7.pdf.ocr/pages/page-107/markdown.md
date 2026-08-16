|  Level | Optional  |
| --- | --- |
|  Interface | Integer  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | ≥0  |

This feature is deprecated (See DeviceEventChannelCount). It indicates the number of message/event channels supported by the device.

### 3.50 DeviceCharacterSet

|  Name | DeviceCharacterSet  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | UTF8 ASCII  |

Character set used by the strings of the device.

Possible values are:

- UTF8: Device use UTF8 character set.
- ASCII: Device use ASCII character set.

### 3.51 DeviceReset

|  Name | DeviceReset  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Recommended  |
|  Interface | ICommand  |
|  Access | Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | -  |