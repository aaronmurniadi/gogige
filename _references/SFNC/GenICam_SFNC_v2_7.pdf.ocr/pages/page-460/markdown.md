|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 24.16 ChunkPartSelector (Deprecated)

|  Name | ChunkPartSelector  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | ≥0  |

This feature is deprecated (See ChunkComponentSelector). It was selecting the individual parts of a multi-component/multi-part buffer to access.

To help backward compatibility, this feature can be included as Invisible in the device's XML.

### 24.17ChunkImage

|  Name | ChunkImage  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IRegister  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Device-specific  |

Returns the entire image data included in the payload.

### 24.18ChunkOffsetX

|  Name | ChunkOffsetX  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |