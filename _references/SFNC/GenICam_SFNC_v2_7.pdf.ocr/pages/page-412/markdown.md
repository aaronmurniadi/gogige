|  Category | Scan3dControl  |
| --- | --- |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read  |
|  Unit | Pixel  |
|  Visibility | Expert  |
|  Values | > 0  |

Returns the focal length of the camera in pixel. The focal length depends on the selected region. The value of this feature takes into account horizontal binning, decimation, or any other function changing the image resolution.

### 21.4.21 Scan3dBaseline

|  Name | Scan3dBaseline  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read  |
|  Unit | m  |
|  Visibility | Expert  |
|  Values | > 0  |

Returns the baseline as the physical distance of two cameras in a stereo camera setup. The value of this feature can be used for 3D reconstruction from disparity images. In this case, the unit of the 3D coordinates corresponds to the unit of the baseline.

### 21.4.22 Scan3dPrincipalPointU

|  Name | Scan3dPrincipalPointU[RegionSelector]  |
| --- | --- |