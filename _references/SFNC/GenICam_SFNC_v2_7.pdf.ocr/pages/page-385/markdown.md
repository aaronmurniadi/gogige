|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

## Calculating World Coordinates from Disparity Images

A disparity image contains the displacement of corresponding pixels between the left and right images of a calibrated and rectified stereo image pair. Since disparity is a displacement in pixel, the Scan3dDistanceUnit of disparity images is always set to Pixel. The distance at a pixel is inversely proportional to the disparity at that pixel. A 3D point cloud can be calculated from a disparity image given additional parameters of the stereo camera, namely the focal length, baseline and image coordinates of the optical center.

The focal length in pixels (i.e. in units of the size of one image pixel after binning and decimation) is given in the feature Scan3dFocalLength. The baseline of the stereo system in the distance unit meter is given in the feature Scan3dBaseline. The optical axis of the camera is the line from the optical center of the camera that orthogonally intersects the image plane. The intersection point is referred to as principal point. This point is typically near the image center. Its horizontal and vertical location is given in the features Scan3dPrincipalPointU and Scan3dPrincipalPointV. These features refer to the selected region and are given relative to the OffsetX and OffsetY of that region.

These parameters of the calibrated stereo camera system permit the 3D reconstruction in the Cartesian camera coordinate system from the disparity D(u,v) at pixel column u and row v:

X(u, v) = (u-Scan3dPrincipalPointU)*Scan3dBaseline/D(u, v)
Y(u, v) = (v-Scan3dPrincipalPointV)*Scan3dBaseline/D(u, v)
Z(u, v) = Scan3dFocalLength*Scan3dBaseline/D(u, v)

These equations permit the reconstruction of all pixels with valid disparities, resulting in a 3D point cloud. A disparity of 0 means that the point is at infinity. This case has to be treated specially when doing 3D reconstruction.

As defined by the equations above, the unit of a reconstructed 3D point is always the same as the unit of the length value given in the Scan3dBaseline parameter.

The pseudo code below shows how to compute a 3D point cloud from a disparity image. The algorithm also works if non-zero OffsetX and OffsetY are used to define a sub-region of the image. This is because Scan3dPrincipalPointU and Scan3dPrincipalPointV, like the column and row indices u and v, are relative to OffsetX and OffsetY. Therefore, no special treatment is necessary in this case.

Scan3dCoordinateSelector = CoordinateC
scaleC = Scan3dCoordinateScale[CoordinateC];
offsetC = Scan3dCoordinateOffset[CoordinateC];
isInvalidC = Scan3dInvalidDataFlag[CoordinateC];
invalidC = Scan3dInvalidDataValue[CoordinateC];

For (row = 0; row < Height; row++)
{
    for (col = 0; col < Width; col++)
    {
    C = imageC[row, col];

    if (isInvalidC == False || C != invalidC)
    {
    D = C*scaleC+offsetC;
    if (D > 0)
    {