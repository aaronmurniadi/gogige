Possible values are:

- RotationX: Rotation around X axis.
- RotationY: Rotation around Y axis.
- RotationZ: Rotation around Z axis.
- TranslationX: Translation along X axis.
- TranslationY: Translation along Y axis.
- TranslationZ: Translation along Z axis.

### 21.4.17 Scan3dTransformValue

|  Name | Scan3dTransformValue[Scan3dExtractionSelector][Scan3dCoordinateTransformSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Specifies the transform value selected. For translations (Scan3dCoordinateTransformSelector = TranslationX/Y/Z) it is expressed in the distance unit of the system, for rotations (Scan3dCoordinateTransformSelector =RotationX/Y/Z) in degrees.

### 21.4.18 Scan3dCoordinateReferenceSelector

|  Name | Scan3dCoordinateReferenceSelector[Scan3dExtractionSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | RotationX RotationY RotationZ  |