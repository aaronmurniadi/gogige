|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Access | Read  |
| --- | --- |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Returns the transform value.

### 24.71 ChunkScan3dCoordinateReferenceSelector

|  Name | ChunkScan3dCoordinateReferenceSelector  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | RotationX RotationY RotationZ TranslationX TranslationY TranslationZ  |

Selector to read a coordinate system reference value defining the transform of a point from one system to the other.

The transformation from Anchor and Transformed to the Reference coordinate system should be implemented as described in section 21.2.2 (Coordinate system position and transformation). Here data describes the current transform, so no selection with Scan3dCoordinateSystemReference is needed.

Possible values are:

- RotationX: Rotation around X axis.
- RotationY: Rotation around Y axis.
- RotationZ: Rotation around Z axis.
- TranslationX: X axis translation.
- TranslationY: Y axis translation.