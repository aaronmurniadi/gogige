|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

- File: The buffer payload contains data of a file. It is used to transfer files, such as JPEG compressed images, which can be stored by the GenTL Producer directly to a hard disk. The user might get a hint how to interpret the buffer by the filename by the "BufferFileName" feature.
- ChunkData: The buffer payload contains chunk data which can be parsed. The chunk data type might be reported through SFNC or deduced from the technology the device is based on. Note that the chunk data can also contain an image. The GenTL Producer should report the presence, position (offset in the buffer) and properties of the image through corresponding BUFFER_INFO_CMD commands.
- JPEG: The buffer payload is a Jpeg formatted image.
- JPEG2000: The buffer payload is a JPEG2000 formatted image.
- H264: The buffer payload is H.264 formatted image data.
- ChunkOnly: The buffer only contains chunk data.
- MultiPart: The buffer payload has multiple parts.
- GenDC: The buffer payload contains a GenDC container.

Corresponds to the BUFFER_INFO_PAYLOADTYPE command of DSGetBufferInfo function.

### 3.5.2.9 BufferNumberOfParts

|  Name | BufferNumberOfParts  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

The number of parts in the current buffer as delivered by the transport mechanism. For non-multipart this is 0, giving that it is not a multipart buffer.

### 3.5.2.10 BufferPartSelector

|  Name | BufferPartSelector  |
| --- | --- |
|  Category | BufferDataInformation  |