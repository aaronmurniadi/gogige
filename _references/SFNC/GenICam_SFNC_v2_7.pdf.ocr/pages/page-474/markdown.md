|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 24.35 ChunkCounterValue

|  Name | ChunkCounterValue[ChunkCounterSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns the value of the selected Chunk counter at the time of the FrameStart event.

### 24.36 ChunkTimerSelector

|  Name | ChunkTimerSelector  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Timer0 (If 0 based), Timer1, Timer2, ...  |

Selects which Timer to retrieve data from.

Possible values are:

- Timer0: Selects the first Timer.
- Timer1: Selects the first Timer.
- Timer2: Selects the second Timer.
• ...