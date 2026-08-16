|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Visibility | Expert  |
| --- | --- |
|  Values | > 0  |

Sets how many Encoder increments/decrements are needed to generate an Encoder output pulse signal.

### 11.2.7 EncoderOutputMode

|  Name | EncoderOutputMode[EncoderSelector]  |
| --- | --- |
|  Category | EncoderControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off PositionUp PositionDown DirectionUp DirectionDown Motion  |

Selects the conditions for the Encoder interface to generate a valid Encoder output signal.

Possible values are:

- Off: No output pulse are generated.
- PositionUp: Output pulses are generated at all new positions in the positive direction. If the encoder reverses no output pulse are generated until it has again passed the position where the reversal started.
- PositionDown: Output pulses are generated at all new positions in the negative direction. If the encoder reverses no output pulse are generated until it has again passed the position where the reversal started.
- DirectionUp: Output pulses are generated at all position increments in the positive direction while ignoring negative direction motion.
- DirectionDown: Output pulses are generated at all position increments in the negative direction while ignoring positive direction motion.
- Motion: Output pulses are generated at all motion increments in both directions.

### 11.2.8 EncoderStatus

|  Name | EncoderStatus[EncoderSelector]  |
| --- | --- |