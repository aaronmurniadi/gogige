On top of the previous standard values, a device might also provide device-specific values.

### 6.10 WhiteClipSelector

|  Name | WhiteClipSelector  |
| --- | --- |
|  Category | AnalogControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | All Red Green Blue Y U V Tap1, Tap2, ...  |

Selects which White Clip to control.

The All White Clip selection is intended to be across all channels or taps, rather than a convenient way to set all the individual channels or tap white clip in a single step. By following this rule, the value read for the All white clip remains valid even when the channel/tap white clips are not all equal.

Possible values are:

- All: White Clip will be applied to all channels or taps.
- Red: White Clip will be applied to the red channel.
- Green: White Clip will be applied to the green channel.
- Blue: White Clip will be applied to the blue channel.
- Y: White Clip will be applied to Y channel.
- U: White Clip will be applied to U channel.
- V: White Clip will be applied to V channel.
- Tap1: White Clip will be applied to Tap 1.
- Tap2: White Clip will be applied to Tap 2.
- ...