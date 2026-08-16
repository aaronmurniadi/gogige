|  Category | EncoderControl  |
| --- | --- |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | EncoderUp EncoderDown EncoderIdle EncoderStatic  |

Returns the motion status of the encoder.

Possible values are:

- EncoderUp: The encoder counter last incremented.
- EncoderDown: The encoder counter last decremented.
- EncoderIdle: The encoder is not active.
- EncoderStatic: No motion within the EncoderTimeout time.

### 11.2.9 EncoderTimeout

|  Name | EncoderTimeout[EncoderSelector]  |
| --- | --- |
|  Category | EncoderControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | us  |
|  Visibility | Expert  |
|  Values | > 0  |

Sets the maximum time interval between encoder counter increments before the status turns to static.