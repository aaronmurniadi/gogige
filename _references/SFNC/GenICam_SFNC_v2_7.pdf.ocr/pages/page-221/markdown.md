|  Category | AnalogControl  |
| --- | --- |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off Once Continuous Device-specific  |

Controls the mode for automatic white balancing between the color channels. The white balancing ratios are automatically adjusted.

Possible values are:

- Off: White balancing is user controlled using BalanceRatioSelector and BalanceRatio.
- Once: White balancing is automatically adjusted once by the device. Once it has converged, it automatically returns to the Off state.
- Continuous: White balancing is constantly adjusted by the device.

On top of the previous standard values, a device might also provide device-specific values.

### 6.15 Gamma

|  Name | Gamma  |
| --- | --- |
|  Category | AnalogControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | >0.0  |

Controls the gamma correction of pixel intensity. This is typically used to compensate for non-linearity of the display system (such as CRT).

Gamma correction is realized by the following formula: