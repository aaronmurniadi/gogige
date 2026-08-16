|  Unit | -  |
| --- | --- |
|  Visibility | Expert  |
|  Values | Sensor Region0 (if 0 based) Region1 Region2 ...  |

Selects which binning engine is controlled by the BinningHorizontal and BinningVertical features.

Possible values are:

- Sensor: Selected features will control the sensor binning.
- Region0: Selected feature will control the region 0 binning.
- Region1: Selected feature will control the region 1 binning.
- Region2: Selected feature will control the region 2 binning.

### 4.25 BinningHorizontalMode

|  Name | BinningHorizontalMode[BinningSelector]  |
| --- | --- |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Sum Average  |

Sets the mode to use to combine horizontal photo-sensitive cells together when BinningHorizontal is used.

Possible values are:

- Sum: The response from the combined cells will be added, resulting in increased sensitivity.
- Average: The response from the combined cells will be averaged, resulting in increased signal/noise ratio.

### 4.26 BinningHorizontal

|  Name | BinningHorizontal[BinningSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IInteger  |