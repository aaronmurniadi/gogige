|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Values | True False  |
| --- | --- |

Flag to indicate that a buffer was filled but an error occurred during that process.

Corresponds to the BUFFER_INFO_IS_INCOMPLETE command of DSGetBufferInfo function.

### 3.5.2.8 BufferPayloadType

|  Name | BufferPayloadType  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Unknown Image RawData File ChunkData JPEG JPEG2000 H264 ChunkOnly MultiPart GenDC  |

Payload type of the data.

- **Unknown**: The GenTL Producer is not aware of the payload type of the data in the provided buffer. For the GenTL Consumer perspective this can be handled as raw data.
- **Image**: The buffer payload contains pure image data. In particular, no chunk data is attached to the image.
- **RawData**: The buffer payload contains raw, unspecified data. For instance, this can be used to send acquisition statistics.