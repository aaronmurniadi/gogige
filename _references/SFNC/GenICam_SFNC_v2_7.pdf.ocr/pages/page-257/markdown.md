### 10.5 Timer features

This section describes the Timers features.

### 10.5.1 TimerSelector

|  Name | TimerSelector  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Timer0 (if 0 based), Timer1, Timer2, ...  |

Selects which Timer to configure.

Possible values are:

- Timer0: Selects the Timer 0.
- Timer1: Selects the Timer 1.
- Timer2: Selects the Timer 2.

### 10.5.2 TimerDuration

|  Name | TimerDuration[TimerSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | us  |
|  Visibility | Expert  |
|  Values | ≥0  |

Sets the duration (in microseconds) of the Timer pulse.

When the Timer reaches the TimerDuration value, a TimerEnd event is generated, the TimerActive signal becomes low and the Timer stops counting until a new trigger happens or it is explicitly reset with TimerReset.