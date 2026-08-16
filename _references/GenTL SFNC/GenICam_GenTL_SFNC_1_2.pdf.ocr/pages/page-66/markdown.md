|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Serial number of the remote device. This value only changes on execution of the DeviceUpdateList command.

Corresponds to the "DeviceSerialNumber" feature of the remote device and is retrieved during device discovery. Note that this feature was added in GenICam SFNC 2.0 and later, thus, for remote devices following an older GenICam SFNC version it corresponds to the "DeviceID" feature of the remote device.

#### 3.2.2.10 DeviceUserID

|  Name | DeviceUserID[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Optional  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

User-programmable device identifier of the remote device. This value only changes on execution of the DeviceUpdateList command.

Corresponds to the "DeviceUserID" feature of the remote device and it is usually retrieved during device discovery.

#### 3.2.2.11 DeviceTLVersionMajor

|  Name | DeviceTLVersionMajor[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |