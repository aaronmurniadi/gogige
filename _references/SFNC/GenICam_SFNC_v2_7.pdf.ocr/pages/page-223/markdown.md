# 7 LUT Control

Features in this chapter describe the Look-up table (LUT) realated features.

## 7.1 LUTControl

|  Name | LUTControl  |
| --- | --- |
|  Category | Root  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that includes the LUT control features.

## 7.2 LUTSelector

|  Name | LUTSelector  |
| --- | --- |
|  Category | LUTControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Luminance Red Green Blue Device-specific  |

Selects which LUT to control.

Possible values are:

- Luminance: Selects the Luminance LUT.
- Red: Selects the Red LUT.