|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

## 6 Analog Control

Features in this chapter describes how to influence the analog features of an image, such as gain, black level, white clip and gamma.

The **Gain**, **BlackLevel** and **Gamma** features will transform the original pixel value Y to a new value Y' according to the following formula:

$$Y' = [(Y + BlackLevel) \cdot Gain]^{Gamma}$$

For some color cameras in Raw or RGB mode, the red/blue channel can be white balanced with respect to the green channel using the Red and blue **BalanceRatio** gain. For cameras in YUV mode the U/V channel can be balanced with respect to the Y channel using the U and V **BalanceRatio**, according to:

$$B' = B(BlueBalanceRatio \cdot Gain)$$

Other color camera controls each color channel gain independently, in which case, the Red, Green and Blue **Gain** features can be used for white balancing.

The automatic functions **GainAuto**, **BlackLevelAuto**, **BalanceWhiteAuto**, **GainAutoTapBalance** and **BlackLevelAutoTapBalance** can be used to auto-adjust a device once or continuously and to turn the function on and off.

Most of the automatic functions have 3 possible values: {**Off**, **Once**, **Continuous**}.

- **Off**: The automatic adjustment is disabled (ie. User control).
- **Once**: The automatic adjustment is performed once by the device. The affected features report the effective values. If necessary, the feature is automatically set to "Off" after the adjustment.
- **Continuous**: The automatic adjustment is continuously done by the device. The affected features report their effective values.

When a device has a specific auto-adjustment capability, it should have a corresponding feature allowing the necessary enumerations.