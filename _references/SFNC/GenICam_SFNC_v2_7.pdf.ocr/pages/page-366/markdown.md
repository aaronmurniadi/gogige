|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

// 3D Scan Region 0: 3D Range as point cloud and Intensity.

RegionSelector = Region0;
RegionMode[Region0] = On;
OffsetX[Region0] = 512;
OffsetY[Region0] = 512;
Width[Region0] = 512;
Height[Region0] = 512;

// Setup Components acquisition for selected 3D Region 0.

ComponentSelector[Region0] = Range;
ComponentEnable[Region0][Range] = True;
PixelFormat[Region0][Range] = Coord3D_ABC8_Planar;
ComponentSelector[Region0] = Intensity;
ComponentEnable[Region0][Intensity] = True;
PixelFormat[Region0][Intensity] = Mono8;

// 3D Scan Region 1: 3D Range as point cloud and Intensity.

RegionSelector = Region1;
RegionMode[Region1] = On;
OffsetX[Region1] = 1024;
OffsetY[Region1] = 0;
Width[Region1] = 400;
Height[Region1] = 256;

// Setup Components acquisition for selected 3D Region 1.

ComponentSelector[Region1] = Range;
ComponentEnable[Region1][Range] = True;
PixelFormat[Region1][Range] = Coord3D_ABC8_Planar;
ComponentSelector[Region0] = Intensity;
ComponentEnable[Region0][Intensity] = True;
PixelFormat[Region0][Intensity] = Mono8;

// Get Scale & offset transforming from transmitted to world coordinates.

// Values in the comments are based on the following example FOV & data use:
// FOV ~25x25x25 inch, X & Y data covers [-12.5,12.5] inch range Z data in [25,50] inch range.
// Invalid data tagged as 0 in all Z coordinate image.
Scan3dCoordinateSelector = CoordinateA;    // CoordinateA -> X in Cartesian system.
scaleA = Scan3dCoordinateScale[CoordinateA];    // e.g. 0.1
offsetA = Scan3dCoordinateOffset[CoordinateA];    // e.g. -12.5
bboxMinA = Scan3dAxisMin[CoordinateA];    // e.g. 1
bboxMaxA = Scan3dAxisMax[CoordinateA];    // e.g. 250

Scan3dCoordinateSelector = CoordinateB;    // CoordinateB -> Y in Cartesian system.
scaleB = Scan3dCoordinateScale[CoordinateB]; // e.g. 0.1
offsetB = Scan3dCoordinateOffset[CoordinateB]; // e.g. -12.5
bboxMinB = Scan3dAxisMin[CoordinateB];    // e.g. 1
bboxMaxB = Scan3dAxisMax[CoordinateB];    // e.g. 250

Scan3dCoordinateSelector = CoordinateC;    // CoordinateC -> Z in Cartesian system.
// Negative scale & large offset to switch Z.
scaleC = Scan3dCoordinateScale[CoordinateC];    // e.g. -0.1
offsetC = Scan3dCoordinateOffset[CoordinateC];    // e.g. 50
if (Scan3dInvalidDataFlag[CoordinateC] == true)
    invalidValueC = Scan3dInvalidDataValue[CoordinateC];    // e.g. = 0
bboxMinC = Scan3dAxisMin[CoordinateC];    // e.g. 1