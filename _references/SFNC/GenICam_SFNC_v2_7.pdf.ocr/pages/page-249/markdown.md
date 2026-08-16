|  Category | CounterAndTimerControl  |
| --- | --- |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | RisingEdge FallingEdge AnyEdge LevelHigh LevelLow  |

Selects the Activation mode of the Counter Reset Source signal.

Possible values are:

- RisingEdge: Resets the counter on the Rising Edge of the signal.
- FallingEdge: Resets the counter on the Falling Edge of the signal.
- AnyEdge: Resets the counter on the Falling or rising Edge of the selected signal.
- LevelHigh: Resets the counter as long as the selected signal level is High.
- LevelLow: Resets the counter as long as the selected signal level is Low.

### 10.3.6 CounterReset

|  Name | CounterReset[CounterSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Does a software reset of the selected Counter and starts it. The counter starts counting events immediately after the reset unless a Counter trigger is active. CounterReset can be used to reset the Counter independently from the CounterResetSource. To disable the counter temporarily, set CounterEventSource to Off.