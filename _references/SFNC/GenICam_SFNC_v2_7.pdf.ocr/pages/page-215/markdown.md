Sets the mode for automatic gain balancing between the sensor color channels or taps. The gain coefficients of each channel or tap are adjusted so they are matched.

Possible values are:

- Off: Gain tap balancing is user controlled using Gain.
- Once: Gain tap balancing is automatically adjusted once by the device. Once it has converged, it automatically returns to the Off state.
- Continuous: Gain tap balancing is constantly adjusted by the device.

On top of the previous standard values, a device might also provide device-specific values.

### 6.6 BlackLevelSelector

|  Name | BlackLevelSelector  |
| --- | --- |
|  Category | AnalogControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | All Red Green Blue Y U V Tap1, Tap2, ...  |

Selects which Black Level is controlled by the various Black Level features.

The All Black Level selection is intended to be across all channels or taps, rather than a convenient way to set all the individual channels or taps black levels in a single step. By following this rule, the value read for the All black level remains valid even when the channel/tap black levels are not all equal.

Possible values are:

- All: Black Level will be applied to all channels or taps.
- Red: Black Level will be applied to the red channel.
- Green: Black Level will be applied to the green channel.
- Blue: Black Level will be applied to the blue channel.