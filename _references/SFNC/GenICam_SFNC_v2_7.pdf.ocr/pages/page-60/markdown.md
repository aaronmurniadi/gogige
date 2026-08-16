![img-6.jpeg](img-6.jpeg)

|  EncoderMode[EncoderSelector] | O | IEnumeration | R/W | - | E | Selects if the count of encoder uses FourPhase mode with jitter filtering or the HighResolution mode without jitter filtering.  |
| --- | --- | --- | --- | --- | --- | --- |
|  EncoderDivider[EncoderSelector] | O | IInteger | R/W | - | E | Sets how many Encoder increments/decrements are needed to generate an Encoder output pulse signal.  |
|  EncoderOutputMode[EncoderSelector] | O | IEnumeration | R/W | - | E | Selects the conditions for the Encoder interface to generate a valid Encoder output signal.  |
|  EncoderStatus[EncoderSelector] | O | IEnumeration | R | - | E | Returns the motion status of the encoder.  |
|  EncoderTimeout[EncoderSelector] | O | IFloat | R/W | us | E | Sets the maximum time interval between encoder counter increments before the status turns to static.  |
|  EncoderResetSource[EncoderSelector] | R | IEnumeration | R/W | - | E | Selects the signals that will be the source to reset the Encoder.  |
|  EncoderResetActivation[EncoderSelector] | R | IEnumeration | R/W | - | E | Selects the Activation mode of the Encoder Reset Source signal.  |
|  EncoderReset[EncoderSelector] | R | ICommand | (R)/W | - | E | Does a software reset of the selected Encoder and starts it.  |
|  EncoderValue[EncoderSelector] | R | IInteger | R/W | - | E | Reads or writes the current value of the position counter of the selected Encoder.  |
|  EncoderValueAtReset[EncoderSelector] | R | IInteger | R | - | E | Reads the value of the of the position counter of the selected Encoder when it was reset by a signal or by an explicit EncoderReset command.  |
|  EncoderResolution[EncoderSelector] | O | IFloat | R/W | mm | E | Defines the resolution of one encoder step.  |

## 2.10 Logic Block Control

Contains the features related to the usage of the logic block (See the Logic Block Control chapter for details).

Table 2-10: Logic Block Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  LogicBlockControl | O | ICategory | R | - | G | Category that contains the Logic Block control features.  |
|  LogicBlockSelector | O | IEnumeration | R/W | - | G | Specifies the Logic Block to configure.  |
|  LogicBlockFunction[LogicBlockSelector] | O | IEnumeration | R/W | - | G | Selects the combinational logic Function of the Logic Block to configure.  |