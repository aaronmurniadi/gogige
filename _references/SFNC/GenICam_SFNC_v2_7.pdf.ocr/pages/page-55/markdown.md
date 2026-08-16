|  **GEN<i>CAM** |   |   |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  MultiSlopeSaturationThreshold[MultiSlopeKneePointSelector] | O | IFloat | R/(W) | % | E | The percentage of the full saturation that is applied at a certain knee-point of a multi-slope exposure.  |
| --- | --- | --- | --- | --- | --- | --- |
|  MultiSlopeIntensityLimit[MultiSlopeKneePointSelector] | O | IFloat | R/(W) | % | E | The relative intensity which divides intensities influenced by different exposure slopes.  |
|  MultiSlopeExposureGradient[MultiSlopeKneePointSelector] | O | IFloat | R/(W) | - | E | The gradient of the additional slope that is defined by this knee-point.  |

## 2.4 Analog Control

Contains the features related to the video signal conditioning in the analog domain (See the Analog Control chapter for details).

Table 2-4: Analog Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  AnalogControl | O | ICategory | R | - | B | Category that contains the Analog control features.  |
|  GainSelector | O | IEnumeration | R/W | - | B | Selects which Gain is controlled by the various Gain features.  |
|  Gain[GainSelector] | O | IFloat | R/W | - | B | Controls the selected gain as an absolute physical value.  |
|  GainAuto[GainSelector] | O | IEnumeration | R/W | - | B | Sets the automatic gain control (AGC) mode.  |
|  GainAutoBalance | O | IEnumeration | R/W | - | B | Sets the mode for automatic gain balancing between the sensor color channels or taps.  |
|  BlackLevelSelector | O | IEnumeration | R/W | - | E | Selects which Black Level is controlled by the various Black Level features.  |
|  BlackLevel[BlackLevelSelector] | O | IFloat | R/W | - | E | Controls the analog black level as an absolute physical value.  |
|  BlackLevelAuto[BlackLevelSelector] | O | IEnumeration | R/W | - | E | Controls the mode for automatic black level adjustment.  |
|  BlackLevelAutoBalance | O | IEnumeration | R/W | - | E | Controls the mode for automatic black level balancing between the sensor color channels or taps.  |
|  WhiteClipSelector | O | IEnumeration | R/W | - | E | Selects which White Clip to control.  |
|  WhiteClip[WhiteClipSelector] | O | IFloat | R/W | - | E | Controls the maximal intensity taken by the video signal before being clipped as an absolute physical value.  |