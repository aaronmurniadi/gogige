|  Access | Read/(Write)  |
| --- | --- |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Big Little  |

Endianness of the registers of the device.

Possible values are:

- Big: Device's registers are big Endian.
- Little: Device's registers are little Endian.

### 3.60 DeviceTemperatureSelector

|  Name | DeviceTemperatureSelector  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Sensor Mainboard Device-specific  |

Selects the location within the device, where the temperature will be measured.

Possible values are:

- Sensor: Temperature of the image sensor of the camera.
- Mainboard: Temperature of the device's mainboard.

### 3.61 DeviceTemperature

|  Name | DeviceTemperature [DeviceTemperatureSelector]  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read  |