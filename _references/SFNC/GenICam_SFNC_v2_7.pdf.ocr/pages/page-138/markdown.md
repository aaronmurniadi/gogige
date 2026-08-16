|  Unit | -  |
| --- | --- |
|  Visibility | Expert  |
|  Values | Discard Average  |

Sets the mode used to reduce the Vertical resolution when **DecimationVertical** is used.

Possible values are:

- **Discard**: The value of every Nth pixel is kept, others are discarded.
- **Average**: The values of a group of N adjacent pixels are averaged.

### 4.32 DecimationVertical

|  Name | DecimationVertical  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |

Vertical sub-sampling of the image. This reduces the vertical resolution (height) of the image by the specified vertical decimation factor.

A value of 1 indicates that the camera performs no vertical decimation.

### 4.33 ReverseX

|  Name | ReverseX  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Recommended  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Flip horizontally the image sent by the device. The Region of interest is applied after the flipping.