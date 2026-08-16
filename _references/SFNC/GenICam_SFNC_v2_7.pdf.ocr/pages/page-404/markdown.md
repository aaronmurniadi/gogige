|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

for(col = 0; col < Width; col++)
{
    coordA [row,col] = col*scaleA+offsetA;
    coordB [row,col] = row*scaleB+offsetB;
    coordC [row,col] = imageC [row,col]*scaleC+offsetC;
}
}

# RectifiedC_Linescan

The transmitted coordinates can be transformed to world coordinates using the following pseudo-code:

for(row = 0; row < Height; row++)
{
    for(col = 0; col < Width; col++)
    {
    coordA [row,col] = col*scaleA+offsetA;
    ChunkScanLineSelector = row;
    // Note: Here encoder wrap-around handling is not illustrated.
    coordB [row,col] = ChunkEncoderValue[row]*scaleB+offsetB;
    coordC [row,col] = imageC [row,col]*scaleC+offsetC;
    }
}

# DisparityC

The transmitted disparity can be scaled and an offset added using the following pseudo-code (see Section 21.2.3 for 3D reconstruction):

for(row = 0; row < Height; row++)
{
    for(col = 0; col < Width; col++)
    {
    coordC [row,col] = imageC [row,col]*scaleC+offsetC;
    }
}

# DisparityC_Linescan

The transmitted B (Y) coordinate can be transformed to world coordinates using the following pseudo-code:

for(row = 0; row < Height; row++)
{
    for(col = 0; col < Width; col++)
    {
    ChunkScanLineSelector = row;
    // Note: Here encoder wrap-around handling is not illustrated.
    coordB [row,col] = ChunkEncoderValue[row]*scaleB+offsetB;
    coordC [row,col] = imageC [row,col]*scaleC+offsetC; // Just scaling to e.g. pixel scale.
    }
}

### 21.4.8 Scan3dCoordinateSystemReference

|  Name | Scan3dCoordinateSystemReference[Scan3dExtractionSelector]  |
| --- | --- |