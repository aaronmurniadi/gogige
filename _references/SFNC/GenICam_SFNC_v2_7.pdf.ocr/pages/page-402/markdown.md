|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

If an encoder is used to trigger the camera at specific displacement intervals this displacement can be set as the scale, and a "non-linescan" mode can be used.

### Use case Examples:

Here we illustrate the calculation of 3D world coordinates for the calibrated formats above.

In the following examples, consider the common definitions of scale and offset:

// CoordinateA - X/Theta.
ChunkScan3dCoordinateSelector = CoordinateA;
scaleA = ChunkScan3dCoordinateScale[CoordinateA];
offsetA = ChunkScan3dCoordinateOffset[CoordinateA];
// CoordinateB - Y / Phi
ChunkScan3dCoordinateSelector = CoordinateB;
scaleB = ChunkScan3dCoordinateScale[CoordinateB];
offsetB = ChunkScan3dCoordinateOffset[CoordinateB];
// CoordinateC - Z / Rho
ChunkScan3dCoordinateSelector = CoordinateC;
scaleC = ChunkScan3dCoordinateScale[CoordinateC];
offsetC = ChunkScan3dCoordinateOffset[CoordinateC];

Then, for each mode

### CalibratedABC_Grid

The transmitted coordinates can be transformed to world coordinates using the following pseudo-code:

for(row = 0; row < Height; row++)
{
    for(col = 0; col < Width; col++)
    {
    coordA [row,col] = imageA [row,col]*scaleA+offsetA;
    coordB [row,col] = imageB [row,col]*scaleB+offsetB;
    coordC [row,col] = imageC [row,col]*scaleC+offsetC;
    }
}

### CalibratedABC_PointCloud

In this format N pixel values in a vector are transmitted each frame. The transmitted coordinates can be transformed to world coordinates using the following pseudo-code:

for(i=1; i<N; i++)
{
    coordA [i] = imageA [i]*scaleA+offsetA;
    coordB [i] = imageB [i]*scaleB+offsetB;
    coordC [i] = imageC [i]*scaleC+offsetC;
}

### CalibratedAC

The transmitted coordinates can be transformed to world coordinates using the following pseudo-code:

for(row = 0; row < Height; row++)
{