- CoordinateB: The second (Y or Phi) coordinate
- CoordinateC: The third (Z or Rho) coordinate.

An example using this to setup scaling and offset of a Cartesian data stream is given below.

// Get Scale & offset, for all 3 coordinates
Scan3dCoordinateSelector = CoordinateA;    // CoordinateA - X
scaleA = Scan3dCoordinateScale[CoordinateA];    // e.g. 0.5
offsetA = Scan3dCoordinateOffset[CoordinateA]; // e.g. -500

Scan3dCoordinateSelector = CoordinateB;    // CoordinateB - Y
scaleB = Scan3dCoordinateScale[CoordinateB];    // e.g. 0.5
offsetB = Scan3dCoordinateOffset[CoordinateB]; // e.g.-500

// Negative scale & large offset to switch Z direction.
Scan3dCoordinateSelector = CoordinateC;    // CoordinateC - Z
scaleC = Scan3dCoordinateScale[CoordinateA];    // e.g. -0.12
offsetC = Scan3dCoordinateOffset[CoordinateA];    // e.g. 5000

### 21.4.10 Scan3dCoordinateScale

|  Name | Scan3dCoordinateScale[Scan3dExtractionSelector][Scan3dCoordinateSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Scale factor when transforming a pixel from relative coordinates to world coordinates.

A negative scale mirrors the axis. For rectified image axes it is the distance between samples in the rectified image along this axis. See example below.

### 21.4.11 Scan3dCoordinateOffset

|  Name | Scan3dCoordinateOffset[Scan3dExtractionSelector][Scan3dCoordinateSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/(Write)  |
|  Unit | -  |