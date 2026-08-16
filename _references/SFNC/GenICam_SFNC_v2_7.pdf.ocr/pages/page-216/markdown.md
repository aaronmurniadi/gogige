- Y: Black Level will be applied to Y channel.
- U: Black Level will be applied to U channel.
- V: Black Level will be applied to V channel.
- Tap1: Black Level will be applied to Tap 1.
- Tap2: Black Level will be applied to Tap 2.
- ...

### 6.7 BlackLevel

|  Name | BlackLevel[BlackLevelSelector]  |
| --- | --- |
|  Category | AnalogControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Device-specific  |

Controls the analog black level as an absolute physical value. This represents a DC offset applied to the video signal.

The unit and values of this feature are specific to the device and must be defined in the GenICam XML device description file.

For color or multi-tap cameras, BlackLevelSelector indicates which channel access.

### 6.8 BlackLevelAuto

|  Name | BlackLevelAuto[BlackLevelSelector]  |
| --- | --- |
|  Category | AnalogControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off Once  |