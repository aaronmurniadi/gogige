|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Visibility | Expert  |
| --- | --- |
|  Values | -  |

Offset when transforming a pixel from relative coordinates to world coordinates.

# Calibration scaling and offset:

The pseudo code shows how to read and use scale and offset in a full 3D calibrated image (scale and offset from chunk data as in examples above).

for(row = 0; row < Height; row++)
{
    for(col = 0; col < Width; col++)
    {
    coordA[row,col] = imageA[row,col]*scaleA+offsetA;
    coordB[row,col] = imageB[row,col]*scaleB+offsetB;
    coordC[row,col] = imageC[row,col]*scaleC+offsetC;
    }
}

# Rectification scaling and offset:

The pseudo code shows how to read and use scale and offset in a rectified image (scale and offset from chunk data as in examples above).

for(row = 0; row < Height; row++)
{
    for(col = 0; col < Width; col++)
    {
    coordA[row,col] = col*scaleA+offsetA;
    coordB[row,col] = row*scaleB+offsetB;
    coordC[row,col] = imageC[row,col]*scaleC+offsetC;
    }
}

# 21.4.12 Scan3dInvalidDataFlag

|  Name | Scan3dInvalidDataFlag[Scan3dExtractionSelector][Scan3dCoordinateSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |