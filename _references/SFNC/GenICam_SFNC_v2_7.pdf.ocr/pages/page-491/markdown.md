|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Values | -  |
| --- | --- |

Returns the Scale for the selected coordinate axis of the image included in the payload.

### 24.64 ChunkScan3dCoordinateOffset

|  Name | ChunkScan3dCoordinateOffset[ChunkScan3dCoordinateSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Returns the Offset for the selected coordinate axis of the image included in the payload.

### 24.65 ChunkScan3dInvalidDataFlag

|  Name | ChunkScan3dInvalidDataFlag[ChunkScan3dCoordinateSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Returns if a specific non-valid data flag is used in the data stream.

Possible values are:

- False: Default value. No specific value identifies missing or invalid points.
- True: The InvalidDataValue specifies a special non-valid value.