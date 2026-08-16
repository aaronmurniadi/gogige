|  Values | ≥0  |
| --- | --- |

Reads or writes the current value (in microseconds) of the selected Timer.

Writing TimerValue is typically used to set the start value.

### 10.5.6 TimerStatus

|  Name | TimerStatus[TimerSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | TimerIdle TimerTriggerWait TimerActive TimerCompleted  |

Returns the current status of the Timer.

Possible values are:

- TimerIdle: The Timer is idle.
- TimerTriggerWait: The Timer is waiting for a start trigger.
- TimerActive: The Timer is counting for the specified duration.
- TimerCompleted: The Timer reached the TimerDuration count.

### 10.5.7 TimerTriggerSource

|  Name | TimerTriggerSource[TimerSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |