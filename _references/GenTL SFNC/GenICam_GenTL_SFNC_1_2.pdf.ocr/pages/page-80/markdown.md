|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Values | Any NULL-terminated string  |
| --- | --- |

Interface-wide unique identifier of this device.

Corresponds to the DEVICE_INFO_ID command of DevGetInfo function.

### 3.3.1.3 DeviceSerialNumber

|  Name | DeviceSerialNumber  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Serial number of the remote device.

Corresponds to the "DeviceSerialNumber" feature of the remote device and usually is retrieved via the bootstrap register of the remote device. Note that this feature has been added in GenICam SFNC 2.0, thus, for remote devices following an older GenICam SFNC version it corresponds to the "DeviceID" feature of the remote device.

Corresponds to the DEVICE_INFO_SERIAL_NUMBER command of DevGetInfo function.

### 3.3.1.4 DeviceUserID

|  Name | DeviceUserID  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Optional  |
|  Interface | IString  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |