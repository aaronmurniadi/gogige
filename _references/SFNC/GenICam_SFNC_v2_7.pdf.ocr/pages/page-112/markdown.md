|  GEN<i>CAM |   | ![img-16.jpeg](img-16.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Unit | C  |
| --- | --- |
|  Visibility | Expert  |
|  Values | Device-specific  |

Device temperature in degrees Celsius (C). It is measured at the location selected by DeviceTemperatureSelector.

3.62 DeviceClockSelector

|  Name | DeviceClockSelector  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | SensorSensorDigitizationCameraLinkDevice-specific  |

Selects the clock frequency to access from the device.

Possible values are:

- Sensor: Clock frequency of the image sensor of the camera.
- SensorDigitization: Clock frequency of the camera A/D conversion stage.
- CameraLink: Frequency of the Camera Link clock.

3.63 DeviceClockFrequency

|  Name | DeviceClockFrequency[DeviceClockSelector]  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/(Write)  |
|  Unit | Hz  |
|  Visibility | Expert  |
|  Values | ≥0  |