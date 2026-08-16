### 24.7 ChunkRegionID

|  Name | ChunkRegionID  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Region0 (if 0 based) Region1 Region2 ... Scan3dExtraction0 (if 0 based) Scan3dExtraction1 Scan3dExtraction2 ...  |

Returns the Identifier of Region that the image comes from.

This generally maps to the corresponding RegionSelector feature.

Possible values are:

- Region0: Image comes from the Region 0.
- Region1: Image comes from the Region 1.
- Region2: Image comes from the Region 2.
- ...
- Scan3dExtraction0: Image comes from the Scan3dExtraction output Region 0.
- Scan3dExtraction1: Image comes from the Scan3dExtraction output Region 1.
- Scan3dExtraction2: Image comes from the Scan3dExtraction output Region 2.
- ...

### 24.8 ChunkRegionIDValue

|  Name | ChunkRegionIDValue[ChunkRegionSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |