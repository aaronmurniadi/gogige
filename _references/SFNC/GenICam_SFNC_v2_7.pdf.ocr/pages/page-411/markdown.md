|   | TranslationX TranslationY TranslationZ  |
| --- | --- |

Sets the index to read a coordinate system reference value defining the transform of a point from the current (Anchor or Transformed) system to the reference system.

The transformation from Anchor and Transformed to the Reference coordinate system should be implemented as described in section 21.2.2 (Coordinate system position and transformation).

Possible values are:

- RotationX: Rotation around X axis.
- RotationY: Rotation around Y axis.
- RotationZ: Rotation around Z axis.
- TranslationX: X axis translation.
- TranslationY: Y axis translation.
- TranslationZ: Z axis translation.

### 21.4.19 Scan3dCoordinateReferenceValue

|  Name | Scan3dCoordinateReferenceValue[Scan3dExtractionSelector] [Scan3dCoordinateReferenceSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Returns the reference value selected. Reads the value of a rotation or translation value for the current (Anchor or Transformed) coordinate system transformation to the Reference system.

### 21.4.20 Scan3dFocalLength

|  Name | Scan3dFocalLength[RegionSelector]  |
| --- | --- |