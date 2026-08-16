### 6.11 WhiteClip

|  Name | WhiteClip[WhiteClipSelector]  |
| --- | --- |
|  Category | AnalogControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Device-specific  |

Controls the maximal intensity taken by the video signal before being clipped as an absolute physical value. The video signal will never exceed the white clipping point: it will saturate at that level.

The unit and values of this feature are specific to the device and must be defined in the GenICam XML device description file.

For color or multi-tap cameras, WhiteClipTapSelector indicates the channel to control.

### 6.12 BalanceRatioSelector

|  Name | BalanceRatioSelector  |
| --- | --- |
|  Category | AnalogControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | All Red Green Blue Y U V Tap1, Tap2, ...  |

Selects which Balance ratio to control.

Possible values are: