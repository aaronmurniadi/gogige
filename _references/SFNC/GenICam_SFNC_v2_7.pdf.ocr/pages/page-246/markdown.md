- LinkTriggerMissed0, LinkTriggerMissed1, LinkTriggerMissed2, ...: Counts the number of Link Triggers missed.
- TimestampTick: Counts the number of clock ticks of the Timestamp clock. Can be used to create a programmable timer.

### 10.3.3 CounterEventActivation

|  Name | CounterEventActivation[CounterSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | RisingEdge FallingEdge AnyEdge  |

Selects the Activation mode Event Source signal.

Possible values are:

- RisingEdge: Counts on the Rising Edge of the signal.
- FallingEdge: Counts on the Falling Edge of the signal.
- AnyEdge: Counts on the Falling or rising Edge of the selected signal.

### 10.3.4 CounterResetSource

|  Name | CounterResetSource[CounterSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off CounterTrigger AcquisitionTrigger AcquisitionTriggerMissed  |