- True: The level of the Line signal is High.
- False: The level of the Line signal is Low.

### 9.2.6 LineStatusAll

|  Name | LineStatusAll  |
| --- | --- |
|  Category | DigitalIOControl  |
|  Level | Optional  |
|  Interface | Integer  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Device-specific  |

Returns the current status of all available Line signals at time of polling in a single bitfield.
The order is Line0 (If 0 based), Line1, Line2,... (lsb to msb).

### 9.2.7 LineSource

|  Name | LineSource[LineSelector]  |
| --- | --- |
|  Category | DigitalIOControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off UserOutput0, UserOutput1, UserOutput2, ... AcquisitionTriggerWait AcquisitionTrigger AcquisitionTriggerMissed AcquisitionActive FrameTriggerWait FrameTrigger FrameTriggerMissed FrameActive ExposureActive LineTriggerWait  |