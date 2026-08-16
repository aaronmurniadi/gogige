Minimum valid transmitted coordinate value of the selected Axis.

The values are for information purposes to e.g. facilitate scaling of display resolution.

### 21.4.15 Scan3dAxisMax

|  Name | Scan3dAxisMax[Scan3dExtractionSelector][Scan3dCoordinateSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Maximum valid transmitted coordinate value of the selected Axis.

The values are for information purposes to e.g. facilitate scaling of display resolution.

### 21.4.16 Scan3dCoordinateTransformSelector

|  Name | Scan3dCoordinateTransformSelector[Scan3dExtractionSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | RotationX RotationY RotationZ TranslationX TranslationY TranslationZ  |

Sets the index to read/write a coordinate transform value.

The transform from Anchor to Transformed coordinate system should be implemented as described in section 21.2.2 (Coordinate system position and transformation).