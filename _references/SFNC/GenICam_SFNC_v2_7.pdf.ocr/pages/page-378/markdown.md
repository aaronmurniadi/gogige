|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

// Calculating World coordinates - Rectified.
// ***
// Reading chunk data from received buffer and scaling Rectified image
// CoordinateA -> X
ChunkScan3dCoordinateSelector = CoordinateA;
scaleA = ChunkScan3dCoordinateScale[Scan3dExtraction0][CoordinateA];
offsetA = ChunkScan3dCoordinateOffset[Scan3dExtraction0][CoordinateA];
// CoordinateB -> Y
ChunkScan3dCoordinateSelector = CoordinateB;
scaleB = ChunkScan3dCoordinateScale[Scan3dExtraction0][CoordinateB];
offsetB = ChunkScan3dCoordinateOffset[Scan3dExtraction0][CoordinateB];
// CoordinateC -> Z
ChunkScan3dCoordinateSelector = CoordinateC;
scaleC = ChunkScan3dCoordinateScale[Scan3dExtraction0][CoordinateC];
offsetC = ChunkScan3dCoordinateOffset[Scan3dExtraction0][CoordinateC];

for (row = 0; row < Height; row++)
{
    for (col = 0; col < Width; col++)
    {
    xCoord [row,col] = col*scaleA+offsetA;
    yCoord [row,col] = row*scaleB+offsetB;
    zCoord [row,col] = imageC[row,col]*scaleC+offsetC;
    }
}

## 21.2 Formatting and interpreting 3D data

The 3d Scan Control relates to what the transmitted 3D data represents, i.e. to formatting and interpreting 3D data from a sensor. Scan3d Control extends the Image Format Control section, and there is a dependency between Scan3d output format and PixelFormat. That is, for a given output format of the 3D data there is a limited number of valid pixel formats in the same way as for a typical 2D camera.

The main parts are

- **Coordinate system:** The 3D data may be expressed in Cartesian, Spherical or Cylindrical coordinates. It is also important where the coordinate system is located relative to the device.
- **Calibration:** The camera may be calibrated to a world coordinate system. The calibration can also include a rectification to a uniform sampling grid. Rectification can give 2.5D range images suitable for standard image processing.
- **Transformation:** The position and pose of the 3D data can be transformed to give output in a coordinate system relevant to the application.

Note that data may also be uncalibrated in which case information on coordinate system and location is meaningless. Therefore many features can be skipped in some cameras, or visible only in some cases depending on e.g. calibration mode.