|  Access | Read/Write  |
| --- | --- |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |

Number of horizontal photo-sensitive cells to combine together. This reduces the horizontal resolution (width) of the image.

A value of 1 indicates that no horizontal binning is performed by the camera.

### 4.27 BinningVerticalMode

|  Name | BinningVerticalMode[BinningSelector]  |
| --- | --- |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Sum Average  |

Sets the mode to use to combine vertical photo-sensitive cells together when BinningVertical is used.

Possible values are:

- Sum: The response from the combined cells will be added, resulting in increased sensitivity.
- Average: The response from the combined cells will be averaged, resulting in increased signal/noise ratio.

### 4.28 BinningVertical

|  Name | BinningVertical[BinningSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |

Number of vertical photo-sensitive cells to combine together. This reduces the vertical resolution (height) of the image.