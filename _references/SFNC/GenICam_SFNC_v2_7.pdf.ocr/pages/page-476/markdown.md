|  Visibility | Expert  |
| --- | --- |
|  Values | Encoder0 (If 0 based), Encoder1, Encoder2, ...  |

Selects which Encoder to retrieve data from.

Possible values are:

- Encoder0: Selects the first Encoder.
- Encoder1: Selects the first Encoder.
- Encoder2: Selects the second Encoder.
- ...

### 24.40 ChunkEncoderValue

|  Name | ChunkEncoderValue[ChunkEncoderSelector][ChunkScanLineSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | Integer  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Returns the counter's value of the selected Encoder at the time of the FrameStart in area scan mode or the counter's value at the time of the LineStart selected by ChunkScanLineSelector in Linescan mode.

### 24.41 ChunkEncoderStatus

|  Name | ChunkEncoderStatus[ChunkEncoderSelector][ChunkScanLineSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | Enumeration  |
|  Access | Read  |
|  Unit | -  |