|  GEN<i>CAM |   | ![img-36.jpeg](img-36.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Category | ImageFormatControl  |
| --- | --- |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | >0  |

Height of the image provided by the device (in pixels).

This reflects the current Region of interest. The maximum value of this feature takes into account vertical binning, decimation, or any other function changing the maximum vertical dimensions of the image and is typically equal to HeightMax - OffsetY.

This feature is generally mandatory for transmitters and transceivers of most Transport Layers.

4.20 OffsetX

|  Name | OffsetX[RegionSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥0  |

Horizontal offset from the origin to the region of interest (in pixels).

4.21 OffsetY

|  Name | OffsetY[RegionSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥0  |