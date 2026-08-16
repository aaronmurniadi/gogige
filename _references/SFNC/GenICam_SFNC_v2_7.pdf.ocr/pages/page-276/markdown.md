Does a software reset of the selected Encoder and starts it. The Encoder starts counting events immediately after the reset. EncoderReset can be used to reset the Encoder independently from the EncoderResetSource.

Note that the value of the Encoder at time of reset is automatically latched and reflected in the EncoderValueAtReset.

### 11.2.13 EncoderValue

|  Name | EncoderValue[EncoderSelector]  |
| --- | --- |
|  Category | EncoderControl  |
|  Level | Recommended  |
|  Interface | Integer  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Reads or writes the current value of the position counter of the selected Encoder.

Writing to EncoderValue is typically used to set the start value of the position counter.

### 11.2.14 EncoderValueAtReset

|  Name | EncoderValueAtReset[EncoderSelector]  |
| --- | --- |
|  Category | EncoderControl  |
|  Level | Recommended  |
|  Interface | Integer  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Reads the value of the of the position counter of the selected Encoder when it was reset by a signal or by an explicit EncoderReset command.

It represents the last Encoder value latched before resetting the Encoder.