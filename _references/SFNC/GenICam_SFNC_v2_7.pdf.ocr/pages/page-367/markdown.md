|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

bboxMaxC = Scan3dAxisMax[CoordinateC];

// e.g. 250

The data output could in this case be formatted as:

|  Component | Part | Source | Region | data type | data format  |
| --- | --- | --- | --- | --- | --- |
|  Range | 0 | 1 | Region 0 | 3D plane of tri planar image | Coord3D_A8  |
|  Range | 1 | 1 | Region 0 | 3D plane of tri planar image | Coord3D_B8  |
|  Range | 2 | 1 | Region 0 | 3D plane of tri planar image | Coord3D_C8  |
|  Intensity | 3 | 1 | Region 0 | 2D image | Mono8  |
|  Range | 4 | 1 | Region 1 | 3D plane of tri planar image | Coord3D_A8  |
|  Range | 5 | 1 | Region 1 | 3D plane of tri planar image | Coord3D_B8  |
|  Range | 6 | 1 | Region 1 | 3D plane of tri planar image | Coord3D_C8  |
|  Intensity | 7 | 1 | Region 1 | 2D image | Mono8  |
|  Chunk data | 8 | - | All components chunk data if enabled (not shown). | (Chunk data) | GenICam Chunk  |

### Using Transformed coordinate system:

The following code shows how to setup a coordinate system transform, and query the transformed system location and orientation.

// Using Transformed coordinate system.
// ***

// Scan3D Control setup of coordinate system information.
// 3D point cloud out and in sensor pixel grid organization.
Scan3dOutputMode = CalibratedABC_Grid;
Scan3dCoordinateSystemReference = Transformed;

// Setup of transform from Anchor location:
// Rotation vector [30,10,-40] degrees.
// Translation vector [1000,200,403] mm.
Scan3dCoordinateTransformSelector = RotationX;
Scan3dTransformValue[RotationX] = 30;
Scan3dCoordinateTransformSelector = RotationY;
Scan3dTransformValue[RotationY] = 10;
Scan3dCoordinateTransformSelector = RotationZ;
Scan3dTransformValue[RotationZ] = -40;
Scan3dCoordinateTransformSelector = TranslationX;
Scan3dTransformValue[TranslationX] = 1000;
Scan3dCoordinateTransformSelector = TranslationY;
Scan3dTransformValue[TranslationY] = 200;