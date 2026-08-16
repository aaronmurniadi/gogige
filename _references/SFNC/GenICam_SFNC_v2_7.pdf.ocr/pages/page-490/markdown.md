|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Returns the Coordinate System Position of the image included in the payload.

Possible values are:

- Anchor: Default value. Original fixed reference. The coordinate system fixed relative the camera reference point marker is used.
- Transformed: Transformed reference system. The transformed coordinate system is used according to the definition in the rotation and translation matrices.

### 24.62 ChunkScan3dCoordinateSelector

|  Name | ChunkScan3dCoordinateSelector  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | CoordinateA CoordinateB CoordinateC  |

Selects which Coordinate to retrieve data from.

Possible values are:

- CoordinateA: The first (X or Theta) coordinate
- CoordinateB: The second (Y or Phi) coordinate
- CoordinateC: The third (Z or Rho) coordinate.

### 24.63 ChunkScan3dCoordinateScale

|  Name | ChunkScan3dCoordinateScale[ChunkScan3dCoordinateSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |