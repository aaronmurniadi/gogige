- AnalogTap2: Analog gain will be applied to Tap 2.
- ...
- DigitalAll: Gain will be applied to all digital channels or taps.
- DigitalRed: Gain will be applied to the red digital channel.
- DigitalGreen: Gain will be applied to the green digital channel.
- DigitalBlue: Gain will be applied to the blue digital channel.
- DigitalY: Gain will be applied to Y digital channel.
- DigitalU: Gain will be applied to U digital channel.
- DigitalV: Gain will be applied to V digital channel.
- DigitalTap1: Digital gain will be applied to Tap 1.
- DigitalTap2: Digital gain will be applied to Tap 2.
- ...

### 6.3 Gain

|  Name | Gain[GainSelector]  |
| --- | --- |
|  Category | AnalogControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Device-specific  |

Controls the selected gain as an absolute physical value. This is an amplification factor applied to the video signal.

The unit and values of this feature are specific to the device and must be defined in the GenICam XML device description file.

For color or multi-tap cameras, GainSelector indicates the color channel or tap to control.

### 6.4 GainAuto

|  Name | GainAuto[GainSelector]  |
| --- | --- |
|  Category | AnalogControl  |