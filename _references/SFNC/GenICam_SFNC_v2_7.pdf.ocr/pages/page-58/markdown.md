|  UserOutputSelector | R | IEnumeration | R/W | - | E | Selects which bit of the User Output register will be set by UserOutputValue.  |
| --- | --- | --- | --- | --- | --- | --- |
|  UserOutputValue[UserOutputSelector] | R | IBoolean | R/W | - | E | Sets the value of the bit selected by UserOutputSelector.  |
|  UserOutputValueAll | O | IInteger | R/W | - | E | Sets the value of all the bits of the User Output register.  |
|  UserOutputValueAllMask | O | IInteger | R/W | - | E | Sets the write mask to apply to the value specified by UserOutputValueAll before writing it in the User Output register.  |

## 2.8 Counter and Timer Control

Contains the features related to the usage of programmable counters and timers (See the Counter and Timer Control chapter for details).

Table 2-8: Counter and Timer Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  CounterAndTimerControl | R | ICategory | R | - | E | Category that contains the Counter and Timer control features.  |
|  CounterSelector | R | IEnumeration | R/W | - | E | Selects which Counter to configure.  |
|  CounterEventSource[CounterSelector] | R | IEnumeration | R/W | - | E | Select the events that will be the source to increment the Counter.  |
|  CounterEventActivation[CounterSelector] | R | IEnumeration | R/W | - | E | Selects the Activation mode Event Source signal.  |
|  CounterResetSource[CounterSelector] | R | IEnumeration | R/W | - | E | Selects the signals that will be the source to reset the Counter.  |
|  CounterResetActivation[CounterSelector] | R | IEnumeration | R/W | - | E | Selects the Activation mode of the Counter Reset Source signal.  |
|  CounterReset[CounterSelector] | R | ICommand | (R)/W | - | E | Does a software reset of the selected Counter and starts it.  |
|  CounterValue[CounterSelector] | R | IInteger | R/W | - | E | Reads or writes the current value of the selected Counter.  |
|  CounterValueAtReset[CounterSelector] | R | IInteger | R | - | E | Reads the value of the selected Counter when it was reset by a trigger or by an explicit CounterReset command.  |
|  CounterDuration[CounterSelector] | R | IInteger | R/W | - | E | Sets the duration (or number of events) before the CounterEnd event is generated.  |
|  CounterStatus[CounterSelector] | R | IEnumeration | R | - | E | Returns the current status of the Counter.  |