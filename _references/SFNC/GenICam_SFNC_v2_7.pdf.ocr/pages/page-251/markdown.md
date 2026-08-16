|  Interface | Integer  |
| --- | --- |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Sets the duration (or number of events) before the CounterEnd event is generated.

When the counter reaches the CounterDuration value, a CounterEnd event is generated, the CounterActive signal becomes inactive and the counter stops counting until a new trigger happens or it is explicitly reset with CounterReset.

### 10.3.10 CounterStatus

|  Name | CounterStatus[CounterSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | CounterIdle CounterTriggerWait CounterActive CounterCompleted CounterOverflow  |

Returns the current status of the Counter.

Possible values are:

- CounterIdle: The counter is idle.
- CounterTriggerWait: The counter is waiting for a start trigger.
- CounterActive: The counter is counting for the specified duration.
- CounterCompleted: The counter reached the CounterDuration count.
- CounterOverflow: The counter reached its maximum possible count.