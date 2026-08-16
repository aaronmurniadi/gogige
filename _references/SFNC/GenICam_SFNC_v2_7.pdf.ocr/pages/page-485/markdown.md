|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Access | Read  |
| --- | --- |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns the unique identifier of the transfer block used to transport the payload.

The block ID is usually defined by the transport layer and repeated in the chunk for convenience.

### 24.54 ChunkTransferStreamID

|  Name | ChunkTransferStreamID  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Stream0 (if 0 based) Stream1 Stream2 Stream3 ...  |

Returns identifier of the stream that generated this block.

Possible values are:

- Stream0: Data comes from Stream0.
- Stream1: Data comes from Stream1.
- Stream2: Data comes from Stream2.
- Stream3: Data comes from Stream3.

Note that the Stream used can be changed using the RegionDestination feature.

### 24.55 ChunkTransferQueueCurrentBlockCount

|  Name | ChunkTransferQueueCurrentBlockCount  |
| --- | --- |