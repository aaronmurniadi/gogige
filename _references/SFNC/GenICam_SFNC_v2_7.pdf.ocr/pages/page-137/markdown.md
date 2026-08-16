A value of 1 indicates that no vertical binning is performed by the camera.

### 4.29 Decimation Horizontal Mode

|  Name | Decimation Horizontal Mode  |
| --- | --- |
|  Level | Optional  |
|  Interface | I Enumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Discard Average  |

Sets the mode used to reduce the horizontal resolution when Decimation Horizontal is used.

Possible values are:

- Discard: The value of every Nth pixel is kept, others are discarded.
- Average: The values of a group of N adjacent pixels are averaged.

### 4.30 Decimation Horizontal

|  Name | Decimation Horizontal  |
| --- | --- |
|  Category | Image Format Control  |
|  Level | Optional  |
|  Interface | I Integer  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |

Horizontal sub-sampling of the image. This reduces the horizontal resolution (width) of the image by the specified horizontal decimation factor.

A value of 1 indicates that the camera performs no horizontal decimation.

### 4.31 Decimation Vertical Mode

|  Name | Decimation Vertical Mode  |
| --- | --- |
|  Level | Optional  |
|  Interface | I Enumeration  |
|  Access | Read/(Write)  |