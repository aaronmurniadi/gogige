|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Unit | -  |
| --- | --- |
|  Visibility | Expert  |
|  Values | Source0 (if 0 based) Source1 Source2 ...  |

Returns the Identifier of Source that the image comes from.

This generally maps to the corresponding SourceSelector feature.

Possible values are:

- Source0: Image comes from the Source 0.
- Source1: Image comes from the Source 1.
- Source2: Image comes from the Source 2.
• ...

### 24.52 ChunkSourceIDValue

|  Name | ChunkSourceIDValue[ChunkSourceSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns the unique integer Identifier value of the Source that the image comes from.

This generally maps to the corresponding SourceIDValue feature.

### 24.53 ChunkTransferBlockID

|  Name | ChunkTransferBlockID  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IInteger  |