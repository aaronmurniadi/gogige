|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- Off: PoCXP is forced off.
- Tripped: The Link has shut down because of an over-current trip.

This feature shall be present only on Receiver Device that controls the PoCXP.

### 27.7.27 CxpFirstLineTriggerWithFrameStart

|  Name | CxpFirstLineTriggerWithFrameStart  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Specifies if a FrameStart trigger also triggers the first LineStart at the same time.