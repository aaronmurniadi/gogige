Version 2.7.1

Standard Features Naming Convention

|  Access | Read  |
| --- | --- |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥0  |

Minor version of the GenCP protocol supported by the device.

3.30 DeviceMaxThroughput

|  Name | DeviceMaxThroughput  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Bps  |
|  Visibility | Expert  |
|  Values | >0  |

Maximum bandwidth of the data that can be streamed out of the device. This can be used to estimate if the physical connection(s) can sustain transfer of free-running images from the camera at its maximum speed.

3.31 DeviceConnectionSelector

|  Name | DeviceConnectionSelector  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥0  |

Selects which Connection of the device to control.

3.32 DeviceConnectionSpeed

|  Name | DeviceConnectionSpeed[DeviceConnectionSelector]  |
| --- | --- |
|  Category | DeviceControl  |