|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Selects the signal which will be the source of the B input of the Encoder.

Possible values are:

- Off: Counter is stopped.
- Line0, Line1, Line2, ...: Encoder Reverse input is taken from the chosen I/O Line.
• ...

#### 11.2.5 EncoderMode

|  Name | EncoderMode[EncoderSelector]  |
| --- | --- |
|  Category | EncoderControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | FourPhase HighResolution  |

Selects if the count of encoder uses FourPhase mode with jitter filtering or the HighResolution mode without jitter filtering.

Possible values are:

- FourPhase: The counter increments or decrements 1 for every full quadrature cycle with jitter filtering.
- HighResolution: The counter increments or decrements every quadrature phase for high resolution counting, but without jitter filtering.

#### 11.2.6 EncoderDivider

|  Name | EncoderDivider[EncoderSelector]  |
| --- | --- |
|  Category | EncoderControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |