|  Category | Scan3dControl  |
| --- | --- |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Anchor Transformed  |

Defines coordinate system reference location.

Defines if the fixed (Anchor) or floating (Transformed) coordinate system is used.

Possible values are:

- Anchor: Default value. Original fixed reference. The coordinate system fixed relative the camera reference point marker is used.
- Transformed: Transformed reference system. The transformed coordinate system is used according to the definition in the rotation and translation matrices.

### 21.4.9 Scan3dCoordinateSelector

|  Name | Scan3dCoordinateSelector[Scan3dExtractionSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | CoordinateA CoordinateB CoordinateC  |

Selects the individual coordinates in the vectors for 3D information/transformation.

This selector is used for all 3D vectors below, independent if they are position or angle coordinates.

Possible values are:

- CoordinateA: The first (X or Theta) coordinate