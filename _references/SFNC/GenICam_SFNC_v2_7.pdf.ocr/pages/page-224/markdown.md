- Green: Selects the Green LUT.
- Blue: Selects the Blue LUT.

This feature is typically not available when only a single LUT is supported.

### 7.3 LUTEnable

|  Name | LUTEnable[LUTSelector]  |
| --- | --- |
|  Category | LUTControl  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Activates the selected LUT.

### 7.4 LUTIndex

|  Name | LUTIndex[LUTSelector]  |
| --- | --- |
|  Category | LUTControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥0  |

Control the index (offset) of the coefficient to access in the selected LUT.

### 7.5 LUTValue

|  Name | LUTValue[LUTSelector][LUTIndex]  |
| --- | --- |
|  Category | LUTControl  |
|  Level | Optional  |