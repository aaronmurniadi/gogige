|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

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
    xCoord[row,col]= imageA[row,col]*scaleA+offsetA;
    yCoord[row,col]= imageB[row,col]*scaleB+offsetB;
    zCoord[row,col]= imageC[row,col]*scaleC+offsetC;
    }
    }
}

if (ScanMode = CalibratedACLineScan)
{
    // Calculating World coordinates - Linescan.
    // ***
    // Reading chunk data from received buffer and scaling 3D image to world coordinates.
    // The image is in Cartesian coordinates.
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
    xCoord[row,col] = imageA[row,col] * scaleA + offsetA;
    ChunkScanLineSelector = row;
    // Note: Here encoder wrap-around handling is not illustrated.
    yCoord[row,col] = ChunkEncoderValue[ChunkScanLineSelector]*scaleB+offsetB;
    zCoord[row,col] = imageC[row,col]*scaleC+offsetC;
    }
    }
}

if (ScanMode == RectifiedC)
{