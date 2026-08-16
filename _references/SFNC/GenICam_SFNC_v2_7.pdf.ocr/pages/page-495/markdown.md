|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- TranslationZ: Z axis translation.

### 24.72 ChunkScan3dCoordinateReferenceValue

|  Name | ChunkScan3dCoordinateReferenceValue[ChunkScan3dCoordinateReferenceSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Returns the value of a position or pose coordinate for the anchor or transformed coordinate systems relative to the reference point.

### 24.73 ChunkScan3dFocalLength

|  Name | ChunkScan3dFocalLength  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read  |
|  Unit | Pixel  |
|  Visibility | Expert  |
|  Values | > 0  |

Returns the focal length of the camera in pixel. The focal length depends on the selected region. The value of this feature takes into account horizontal binning, decimation, or any other function changing the image resolution.