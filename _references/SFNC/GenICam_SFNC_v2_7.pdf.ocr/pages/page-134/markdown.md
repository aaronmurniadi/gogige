|  GEN<i>CAM |   | ![img-37.jpeg](img-37.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Vertical offset from the origin to the region of interest (in pixels).

4.22 LinePitchEnable

|  Name | LinePitchEnable[RegionSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Recommended  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | TrueFalse  |

This feature controls whether the LinePitch feature is writable. Otherwise LinePitch is implicitly controlled by the combination of features like Width, PixelFormat, etc...

4.23 LinePitch

|  Name | LinePitch[RegionSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | B  |
|  Visibility | Expert  |
|  Values | ≥0  |

Total number of bytes between the starts of 2 consecutive lines. This feature is used to facilitate alignment of image data.

This might be useful if the system has specific limitations, such as having the lines aligned on 32-bit boundaries.

4.24 BinningSelector

|  Name | BinningSelector  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |