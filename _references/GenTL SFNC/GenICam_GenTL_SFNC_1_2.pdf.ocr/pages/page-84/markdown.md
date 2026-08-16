|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Note that these values already follow the updated value list of the "DeviceTLType" feature from GenICam SFNC 2.3. Depending on this value, the transport layer specific features for the chosen transport layer standard have to be considered.

- CameraLink: Camera Link
- CameraLinkHS: Camera Link High Speed
- CoaXPress: CoaXPress
- GigEVision: GigE Vision
- USB3Vision: USB3 Vision
- Custom: Custom transport layer

Corresponds to the DEVICE_INFO_TLTYPE command of DevGetInfo function.

##### 3.3.1.11 DeviceDisplayName

|  Name | DeviceDisplayName  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

User readable name of the device. If this is not defined in the device this should be "VENDOR MODEL (ID)".

Corresponds to the DEVICE_INFO_DISPLAYNAME command of DevGetInfo function.

##### 3.3.1.12 DeviceTimestampFrequency

|  Name | DeviceTimestampFrequency  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |