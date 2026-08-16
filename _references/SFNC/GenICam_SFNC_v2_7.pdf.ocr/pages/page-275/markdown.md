### 11.2.11 EncoderResetActivation

|  Name | EncoderResetActivation[EncoderSelector]  |
| --- | --- |
|  Category | EncoderControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | RisingEdge FallingEdge AnyEdge LevelHigh LevelLow  |

Selects the Activation mode of the Encoder Reset Source signal.

Possible values are:

- RisingEdge: Resets the Encoder on the Rising Edge of the signal.
- FallingEdge: Resets the Encoder on the Falling Edge of the signal.
- AnyEdge: Resets the Encoder on the Falling or rising Edge of the selected signal.
- LevelHigh: Resets the Encoder as long as the selected signal level is High.
- LevelLow: Resets the Encoder as long as the selected signal level is Low.

### 11.2.12 EncoderReset

|  Name | EncoderReset[EncoderSelector]  |
| --- | --- |
|  Category | EncoderControl  |
|  Level | Recommended  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |