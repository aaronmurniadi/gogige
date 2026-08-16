|  GEN<i>CAM |   | ![img-16.jpeg](img-16.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | DeviceEnumeration  |
| --- | --- |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Name of the device vendor. This value only changes on execution of the DeviceUpdateList command.

Corresponds to the "DeviceVendorName" feature of the remote device and is retrieved during device discovery.

##### 3.2.2.7 DeviceModelName

|  Name | DeviceModelName[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Name of the device model. This value only changes on execution of the DeviceUpdateList command.

Corresponds to the "DeviceModelName" feature of the remote device and is retrieved during device discovery.

##### 3.2.2.8 DeviceAccessStatus

|  Name | DeviceAccessStatus[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |