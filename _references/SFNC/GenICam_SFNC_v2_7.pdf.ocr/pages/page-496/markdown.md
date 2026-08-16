|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 24.74 ChunkScan3dBaseline

|  Name | ChunkScan3dBaseline  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read  |
|  Unit | m  |
|  Visibility | Expert  |
|  Values | > 0  |

Returns the baseline as the physical distance of two cameras in a stereo camera setup. The value of this feature can be used for 3D reconstruction from disparity images. In this case, the unit of the 3D coordinates corresponds to the unit of the baseline.

### 24.75 ChunkScan3dPrincipalPointU

|  Name | ChunkScan3dPrincipalPointU  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read  |
|  Unit | Pixel  |
|  Visibility | Expert  |
|  Values | -  |

Returns the value of this feature gives the horizontal position of the principal point, relative to the region origin, i.e. OffsetX. The value of this feature takes into account horizontal binning, decimation, or any other function changing the image resolution.