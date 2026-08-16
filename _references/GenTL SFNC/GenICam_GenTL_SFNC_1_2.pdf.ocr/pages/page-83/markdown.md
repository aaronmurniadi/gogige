|  GEN<i>CAM |   | ![img-25.jpeg](img-25.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

3.3.1.9 DeviceManufacturerInfo

|  Name | DeviceManufacturerInfo  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Manufacturer information about the remote device.

Corresponds to the "DeviceManufacturerInfo" feature of the remote device and is usually retrieved via the bootstrap register of the remote device.

3.3.1.10 DeviceType

|  Name | DeviceType  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | GigEVisionCameraLinkCameraLinkHSCoaXPressUSB3VisionCustom  |

Transport layer type of the device.