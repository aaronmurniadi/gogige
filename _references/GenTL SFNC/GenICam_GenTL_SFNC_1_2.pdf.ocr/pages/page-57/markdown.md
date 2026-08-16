|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Interface | IString  |
| --- | --- |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

A user-friendly name of the Interface.

Corresponds to the INTERFACE_INFO_DISPLAYNAME command of IFGetInfo function.

##### 3.2.1.4 InterfaceType

|  Name | InterfaceType  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | GigEVisionCameraLinkCameraLinkHSCoaXPressUSB3VisionCustom  |

Transport layer type of the interface.

Note that these values already follow the updated value list of the "DeviceTLType" feature from GenICam SFNC 2.3. Depending on this value, the transport layer specific features for the chosen transport layer standard have to be considered.

• CameraLink: Camera Link
• CameraLinkHS: Camera Link High Speed
- CoaXPress: CoaXPress