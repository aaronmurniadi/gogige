- Eight: 8 taps.
- Ten: 10 taps.

### 4.5 WidthMax

|  Name | WidthMax  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |

Maximum width of the image (in pixels). The dimension is calculated after horizontal binning, decimation or any other function changing the horizontal dimension of the image.

WidthMax does not take into account the current Region of interest (Width or OffsetX). Its value must be greater than 0 and less than or equal to SensorWidth (unless an oversampling feature is present).

### 4.6 HeightMax

|  Name | HeightMax  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |

Maximum height of the image (in pixels). This dimension is calculated after vertical binning, decimation or any other function changing the vertical dimension of the image.

HeightMax does not take into account the current Region of interest (Height or OffsetY). Its value must be greater than 0 and for area scan devices, less than or equal to SensorHeight (unless an oversampling feature is present).

### 4.7 RegionSelector

|  Name | RegionSelector  |
| --- | --- |
|  Category | ImageFormatControl  |