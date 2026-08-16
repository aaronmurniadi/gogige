|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Enables the definition of a non-valid flag value in the data stream. Note that the confidence output is an alternate recommended way to identify non-valid pixels. Using a Scan3dInvalidDataValue may give processing penalties due to special handling.

Possible values are:

- False: Default value. No specific value identifies missing or invalid points.
- True: The Scan3dInvalidDataValue specifies a special non-valid value.

### 21.4.13 Scan3dInvalidDataValue

|  Name | Scan3dInvalidDataValue[Scan3dExtractionSelector][Scan3dCoordinateSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Value which identifies a non-valid pixel if Scan3dInvalidDataFlag is enabled.

Typically the invalid data is flagged in one coordinate (Z/Rho) only, but it can be applied to each coordinate.

If the pixel format is integer the value must be mapped to (rounded to) an integer register in the device. Using a floating point NaN during pixel data processing might incur performance penalties, it might be desirable to avoid such values within pixel data whenever possible.

### 21.4.14 Scan3dAxisMin

|  Name | Scan3dAxisMin[Scan3dExtractionSelector][Scan3dCoordinateSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |