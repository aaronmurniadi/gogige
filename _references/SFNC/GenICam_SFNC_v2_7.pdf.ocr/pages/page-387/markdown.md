|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

ScaleHorizontal = DisparityScaleHorizontal/IntensityScaleHorizontal;
ScaleVertical = DisparityScaleVertical/IntensityScaleVertical;

// Generates the Intensity texture image corresponding to the Disparity image.
for (row = 0; row < DisparityHeight; row++)
{
    for (col = 0; col < DisparityWidth; col++)
    {
        IntensityTexturePixel[row][col] = IntensityPixel[row*ScaleHorizontal][col*ScaleVertical];
    }
}

// ...