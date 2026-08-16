### 10.5.3 TimerDelay

|  Name | TimerDelay[TimerSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | us  |
|  Visibility | Expert  |
|  Values | ≥0  |

Sets the duration (in microseconds) of the delay to apply at the reception of a trigger before starting the Timer.

### 10.5.4 TimerReset

|  Name | TimerReset[TimerSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Does a software reset of the selected timer and starts it. The timer starts immediately after the reset unless a timer trigger is active.

### 10.5.5 TimerValue

|  Name | TimerValue[TimerSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | us  |
|  Visibility | Expert  |