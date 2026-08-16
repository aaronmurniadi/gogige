Note that the value of the Counter at time of reset is automatically latched and reflected in the CounterValueAtReset.

### 10.3.7 CounterValue

|  Name | CounterValue[CounterSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Reads or writes the current value of the selected Counter.

Writing to CounterValue is typically used to set the start value.

### 10.3.8 CounterValueAtReset

|  Name | CounterValueAtReset[CounterSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Reads the value of the selected Counter when it was reset by a trigger or by an explicit CounterReset command.

It represents the last counter value latched before resetting the counter.

### 10.3.9 CounterDuration

|  Name | CounterDuration[CounterSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |