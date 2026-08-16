|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Returns the last Timestamp latched with the TimestampLatch command.

### 24.33 ChunkLineStatusAll

|  Name | ChunkLineStatusAll  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns the status of all the I/O lines at the time of the FrameStart internal event.

### 24.34 ChunkCounterSelector

|  Name | ChunkCounterSelector  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Counter0 (If 0 based), Counter1, Counter2,...  |

Selects which counter to retrieve data from.

Possible values are:

- Counter0: Selects the counter 0.
- Counter1: Selects the counter 1.
- Counter2: Selects the counter 2.
• ...