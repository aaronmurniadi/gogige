|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Values | -  |
| --- | --- |

Returns the Maximum Axis value for the selected coordinate axis of the image included in the payload.

### 24.69 ChunkScan3dCoordinateTransformSelector

|  Name | ChunkScan3dCoordinateTransformSelector  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | RotationX RotationY RotationZ TranslationX TranslationY TranslationZ  |

Selector for transform values.

Possible values are:

- RotationX: Rotation around X axis.
- RotationY: Rotation around Y axis.
- RotationZ: Rotation around Z axis.
- TranslationX: Translation along X axis.
- TranslationY: Translation along Y axis.
- TranslationZ: Translation along Z axis.

### 24.70 ChunkScan3dTransformValue

|  Name | ChunkScan3dTransformValue[ChunkScan3dCoordinateTransformSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IFloat  |