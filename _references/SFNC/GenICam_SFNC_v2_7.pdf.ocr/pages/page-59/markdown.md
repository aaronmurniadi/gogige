![img-5.jpeg](img-5.jpeg)

|  CounterTriggerSource[CounterSelector] | R | IEnumeration | R/W | - | E | Selects the source to start the Counter.  |
| --- | --- | --- | --- | --- | --- | --- |
|  CounterTriggerActivation[CounterSelector] | R | IEnumeration | R/W | - | E | Selects the activation mode of the trigger to start the Counter.  |
|  TimerSelector | R | IEnumeration | R/W | - | E | Selects which Timer to configure.  |
|  TimerDuration[TimerSelector] | R | IFloat | R/W | us | E | Sets the duration (in microseconds) of the Timer pulse.  |
|  TimerDelay[TimerSelector] | R | IFloat | R/W | us | E | Sets the duration (in microseconds) of the delay to apply at the reception of a trigger before starting the Timer.  |
|  TimerReset[TimerSelector] | R | ICommand | (R)/W | - | E | Does a software reset of the selected timer and starts it.  |
|  TimerValue[TimerSelector] | R | IFloat | R/W | us | E | Reads or writes the current value (in microseconds) of the selected Timer.  |
|  TimerStatus[TimerSelector] | R | IEnumeration | R | - | E | Returns the current status of the Timer.  |
|  TimerTriggerSource[TimerSelector] | R | IEnumeration | R/W | - | E | Selects the source of the trigger to start the Timer.  |
|  TimerTriggerActivation[TimerSelector] | R | IEnumeration | R/W | - | E | Selects the activation mode of the trigger to start the Timer.  |
|  TimerTriggerArmDelay[TimerSelector] | R | IFloat | R/W | us | E | Sets the minimum period between two valid timer triggers.  |

## 2.9 Encoder Control

Contains the features related to the usage of quadrature encoders (See the Encoder Control chapter for details).

Table 2-9: Quadrature Encoder Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  EncoderControl | O | ICategory | R | - | B | Category that contains the quadrature Encoder Control features.  |
|  EncoderSelector | O | IEnumeration | R/W | - | E | Selects which Encoder to configure.  |
|  EncoderSourceA[EncoderSelector] | O | IEnumeration | R/W | - | E | Selects the signal which will be the source of the A input of the Encoder.  |
|  EncoderSourceB[EncoderSelector] | O | IEnumeration | R/W | - | E | Selects the signal which will be the source of the B input of the Encoder.  |