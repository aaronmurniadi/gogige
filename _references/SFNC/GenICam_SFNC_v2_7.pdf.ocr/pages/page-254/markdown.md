- LogicBlock0 (if 0 based), LogicBlock1, LogicBlock2, ...: Starts with the reception of the Logic Block output signal.
- SoftwareSignal0, SoftwareSignal1, SoftwareSignal2, ...: Starts on the reception of the Software Signal.
- Action0, Action1, Action2, ...: Starts with the assertion of the chosen action signal.
- LinkTrigger0, LinkTrigger1, LinkTrigger2, ...: Starts with the reception of the chosen Link Trigger signal.

### 10.3.12 CounterTriggerActivation

|  Name | CounterTriggerActivation[CounterSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | RisingEdge FallingEdge AnyEdge LevelHigh LevelLow  |

Selects the activation mode of the trigger to start the Counter.

Possible values are:

- RisingEdge: Starts counting on the Rising Edge of the selected trigger signal.
- FallingEdge: Starts counting on the Falling Edge of the selected trigger signal.
- AnyEdge: Starts counting on the Falling or rising Edge of the selected trigger signal.
- LevelHigh: Counts as long as the selected trigger signal level is High.
- LevelLow: Counts as long as the selected trigger signal level is Low.

### 10.4 Timer usage model

This section describes the Timers usage model.