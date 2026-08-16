|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Recommended behavior of the identifier: It should start at a certain minimum value and keep incrementing by one for each frame up to a maximum, then it wraps to the minimum again. Each streaming channel should maintain the Frame ID separately.

Note: For GigE Vision, this chunk is not necessarily the block_id field included in the GVSP headers but can be equal to it.

### 24.50 ChunkSourceSelector

|  Name | ChunkSourceSelector  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Source0 (if 0 based) Source1 Source2 ...  |

Selects which Source to retrieve data from.

This generally maps to the corresponding SourceSelector feature.

Possible values are:

- Source0: Image comes from the Source 0.
- Source1: Image comes from the Source 1.
- Source2: Image comes from the Source 2.
• ...

### 24.51 ChunkSourceID

|  Name | ChunkSourceID  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |