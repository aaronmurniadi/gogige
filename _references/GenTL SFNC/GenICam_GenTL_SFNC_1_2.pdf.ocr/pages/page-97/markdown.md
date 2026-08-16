|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.4.1.3 StreamType

|  Name | StreamType  |
| --- | --- |
|  Category | StreamInformation  |
|  Level | Mandatory  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | GigEVision CameraLink CameraLinkHS CoaXPress USB3Vision Custom  |

Transport layer type of the Data Stream.

Note that these values already follow the updated value list of the "DeviceTLType" feature from GenICam SFNC 2.3. Depending on this value, the transport layer specific features for the chosen transport layer standard have to be considered.

- CameraLink: Camera Link
- CameraLinkHS: Camera Link High Speed
- CoaXPress: CoaXPress
- GigEVision: GigE Vision
- USB3Vision: USB3 Vision
- Custom: Custom transport layer

Corresponds to the STREAM_INFO_TLTYPE command of DSGetInfo function.