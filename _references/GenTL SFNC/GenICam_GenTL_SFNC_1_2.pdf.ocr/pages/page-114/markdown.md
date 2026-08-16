|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|   | Custom  |
| --- | --- |

Transport layer type of the buffer.

Note that these values already follow the updated value list of the "DeviceTLType" feature from GenICam SFNC 2.3. Depending on this value, the transport layer specific features for the chosen transport layer standard have to be considered.

- CameraLink: Camera Link
- CameraLinkHS: Camera Link High Speed
- CoaXPress: CoaXPress
- GigEVision: GigE Vision
- USB3Vision: USB3 Vision
- Custom: Custom transport layer

Corresponds to the BUF_INFO_TLTYPE command of DSGetBufferInfo function.

### 3.5.1.4 BufferSize

|  Name | BufferSize  |
| --- | --- |
|  Category | BufferInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Byte  |
|  Visibility | Expert  |
|  Values | ≥0  |

Size of the buffer in bytes.

Corresponds to the BUF_INFO_SIZE command of DSGetBufferInfo function.

### 3.5.2 Buffer Data Information

Features in this section provide information about the currently filled data in the buffers. Note that for multipart buffers the BufferPartSelector is used to extract information for each part of the buffer.