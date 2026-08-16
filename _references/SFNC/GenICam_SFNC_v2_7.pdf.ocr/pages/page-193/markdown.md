- LevelHigh: Specifies that the trigger is considered valid as long as the level of the source signal is high.
- LevelLow: Specifies that the trigger is considered valid as long as the level of the source signal is low.

### 5.6.6 TriggerOverlap

|  Name | TriggerOverlap[TriggerSelector]  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off ReadOut PreviousFrame PreviousLine  |

Specifies the type trigger overlap permitted with the previous frame or line. This defines when a valid trigger will be accepted (or latched) for a new frame or a new line.

Possible values are:

- Off: No trigger overlap is permitted.
- ReadOut: Trigger is accepted immediately after the exposure period.
- PreviousFrame: Trigger is accepted (latched) at any time during the capture of the previous frame.
- PreviousLine: Trigger is accepted (latched) at any time during the capture of the previous line.

### 5.6.7 TriggerDelay

|  Name | TriggerDelay[TriggerSelector]  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | us  |
|  Visibility | Expert  |
|  Values | Device-specific  |