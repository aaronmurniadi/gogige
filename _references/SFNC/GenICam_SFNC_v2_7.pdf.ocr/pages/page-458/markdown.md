|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Visibility | Expert  |
| --- | --- |
|  Values | Group0, Group1, Group2, ...  |

Selects the component Group from which to retrieve data from.

This generally maps to the corresponding GroupSelector feature.

Possible values are:

- Group0: Selects Components group 0.
- Group1: Selects Components group 1.
- Group2: Selects Components group 2.

### 24.13 ChunkGroupID

|  Name | ChunkGroupID[ChunkGroupSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Group0, Group1, Group2, ...  |

Returns a unique Identifier corresponding to the selected Group of components. This can be used to identify the component Group of a multi-group payload.

This generally maps to the corresponding GroupSelector feature.

Possible values are:

- Group0: Selects Components group 0.
- Group1: Selects Components group 1.
- Group2: Selects Components group 2.

### 24.14 ChunkGroupIDValue

|  Name | ChunkGroupIDValue[ChunkGroupSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |