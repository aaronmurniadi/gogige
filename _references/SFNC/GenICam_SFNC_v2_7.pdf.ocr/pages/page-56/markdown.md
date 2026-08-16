|  BalanceRatioSelector | O | IEnumeration | R/W | - | E | Selects which Balance ratio to control.  |
| --- | --- | --- | --- | --- | --- | --- |
|  BalanceRatio[BalanceRatioSelector] | O | IFloat | R/W | - | E | Controls ratio of the selected color component to a reference color component.  |
|  BalanceWhiteAuto | O | IEnumeration | R/W | - | E | Controls the mode for automatic white balancing between the color channels.  |
|  Gamma | O | IFloat | R/W | - | B | Controls the gamma correction of pixel intensity.  |

## 2.5 LUT Control

Contains the features related to the look-up table (LUT) control (See the LUT Control chapter for details).

Table 2-5: Lut Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  LUTControl | O | ICategory | R | - | E | Category that includes the LUT control features.  |
|  LUTSelector | O | IEnumeration | R/W | - | E | Selects which LUT to control.  |
|  LUTEnable[LUTSelector] | O | IBoolean | R/W | - | E | Activates the selected LUT.  |
|  LUTIndex[LUTSelector] | O | IInteger | R/W | - | G | Control the index (offset) of the coefficient to access in the selected LUT.  |
|  LUTValue[LUTSelector][LUTIndex] | O | IInteger | R/W | - | G | Returns the Value at entry LUTIndex of the LUT selected by LUTSelector.  |
|  LUTValueAll[LUTSelector] | O | IRegister | R/W | - | G | Accesses all the LUT coefficients in a single access without using individual LUTIndex.  |

## 2.6 Color Transformation Control

Contains the features related to the control of the color transformation (See the Color Transformation Control chapter for details).

Table 2-6: Color Transformation summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |