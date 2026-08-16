|  Unit | um  |
| --- | --- |
|  Visibility | Guru  |
|  Values | >0  |

Physical size (pitch) in the y direction of a photo sensitive pixel unit.

### 4.1 SensorName

|  Name | SensorName  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Any NULL-terminated string  |

Product name of the imaging Sensor.

### 4.2 SensorShutterMode

|  Name | SensorShutterMode  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Global Rolling GlobalReset  |

Specifies the shutter mode of the device.

Possible values are:

- Global: The shutter opens and closes at the same time for all pixels. All the pixels are exposed for the same length of time at the same time.