|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

for(col = 0; col < Width; col++)
{
    coordA [row,col] = imageA [row,col]*scaleA+offsetA;
    coordB [row,col] = row*scaleB+offsetB;
    coordC [row,col] = imageC [row,col]*scaleC+offsetC;
}
}

#### CalibratedAC_Linescan

The transmitted coordinates can be transformed to world coordinates using the following pseudo-code:

for(row = 0; row < Height; row++)
{
    for(col = 0; col < Width; col++)
    {
    coordA [row,col] = imageA [row,col]*scaleA+offsetA;
    ChunkScanLineSelector = row;
    // Note: Here encoder wrap-around handling is not illustrated.
    coordB [row,col] = ChunkEncoderValue[row]*scaleB+offsetB;
    coordC [row,col] = imageC [row,col]*scaleC+offsetC;
    }
}

#### CalibratedC

The transmitted coordinates can be transformed to world coordinates using the following pseudo-code:

for(row = 0; row < Height; row++)
{
    for(col = 0; col < Width; col++)
    {
    coordC [row,col] = imageC [row,col]*scaleC+offsetC;
    }
}

#### CalibratedC_Linescan

The transmitted coordinates can be transformed to world coordinates using the following pseudo-code:

for(row = 0; row < Height; row++)
{
    for(col = 0; col < Width; col++)
    {
    ChunkScanLineSelector = row;
    // Note: Here encoder wrap-around handling is not illustrated.
    coordB [row,col] = ChunkEncoderValue[row]*scaleB+offsetB;
    coordC [row,col] = imageC [row,col]*scaleC+offsetC;
    }
}

#### RectifiedC

The transmitted coordinates can be transformed to world coordinates using the following pseudo-code:

for(row = 0; row < Height; row++)
{