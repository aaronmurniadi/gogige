|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

X = (col-Scan3dPrincipalPointU)*Scan3dBaseline/D;
Y = (row-Scan3dPrincipalPointV)*Scan3dBaseline/D;
Z = Scan3dFocalLength*Scan3dBaseline/D;
// do something with the reconstructed point X, Y, Z
}
else
{
    // point is reconstructed at infinity
}
}
else
{
    // value is marked as invalid
}
}

### 21.2.4 Mapping Disparity and Intensity of Different Resolutions

For stereo cameras it is common to provide the disparity image in a lower resolution than the camera intensity image. In that case, the width and height of the components will differ, even if they come from the same source and region. The difference in resolution in that case will be given by the ratio of binning and / or decimation factors between those components.

If the 3D reconstruction includes texturing, it becomes necessary to adapt the intensity images resolution to the disparity. It is also important to note that if the components provide images in different resolutions, then all parameters that refer to pixel positions depend on the component selector including the location of the principle point and focal length.

The following code snippet demonstrates the main calculations that are necessary in that case:

// Mapping intensity and disparity images coming from same source and region but that differ
// in scale as expressed by binning and/or decimation features.

// ...

// Read Disparity scale.
ComponentSelector = Disparity;
DisparityScaleHorizontal = BinningHorizontal[Source0][Region0][Disparity];
DisparityScaleHorizontal *= DecimationHorizontal[Source0][Region0][Disparity];
DisparityScaleVertical = BinningDecimationVertical[Source0][Region0][Disparity];
DisparityScaleVertical *= DecimationVertical[Source0][Region0][Disparity];

// Read Intensity scale.
ComponentSelector = Intensity;
IntensityScaleHorizontal = BinningHorizontal[Source0][Region0][Intensity];
IntensityScaleHorizontal *= DecimationHorizontal[Source0][Region0][Intensity];
IntensityScaleVertical = BinningVertical[Source0][Region0][Intensity];
IntensityScaleVertical *= DecimationVertical[Source0][Region0][Intensity];

// Calculates the Intensity to Disparity scale ratio.