|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|   | RegionIDValue ComponentID ComponentIDValue GroupIDValue TransferBlockID TransferStreamID TransferQueueCurrentBlockCount StreamChannelID SequencerSetActive Scan3dDistanceUnit Scan3dOutputMode Scan3dCoordinateSystem Scan3dCoordinateSystemReference Scan3dCoordinateScale Scan3dCoordinateOffset Scan3dInvalidDataFlag Scan3dInvalidDataValue Scan3dAxisMin Scan3dAxisMax Scan3dCoordinateTransformValue Scan3dCoordinateReferenceValue  |
| --- | --- |

Selects which Chunk to enable or control.

### 24.5 ChunkEnable

|  Name | ChunkEnable[ChunkSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Enables the inclusion of the selected Chunk data in the payload of the image.