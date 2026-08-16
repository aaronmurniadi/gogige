|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Offset of the image data from the beginning of the delivered buffer in bytes. Applies for example when delivering the image as part of chunk data or on technologies requiring specific buffer alignment.

Corresponds to the BUFFER_INFO_IMAGEOFFSET command of DSGetBufferInfo function.

3.5.2.25 BufferPixelFormat

|  Name | BufferPixelFormat[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Mono1p Mono2p Mono4p Mono8 Mono8s Mono10 Mono10c3a64 Mono10c3p32 Mono10g12 Mono10msb Mono10p Mono10pmsb Mono10s Mono12 Mono12g Mono12msb Mono14 Mono16 R8 G8 B8 RGB8 RGB8_Planar  |