TimerSelector         = Timer1;
TimerDuration         = 200;
TimerTriggerSource     = Action3;

LineSelector           = Line2;
LineMode               = Output;
LineSource             = Timer1Active;

ActionDeviceKey        = 0x12345678;
ActionSelector         = 3;
ActionGroupKey          = 0x1;
ActionGroupMask         = 0x7;

// Here the Device is ready to receive the Action Command
// from an external source.

## 14.2 Action Control Features

This section describes the action control features.

### 14.2.1 ActionControl

|  Name | ActionControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | -  |

Category that contains the Action control features.

### 14.2.2 ActionUnconditionalMode

|  Name | ActionUnconditionalMode  |
| --- | --- |
|  Category | ActionControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |