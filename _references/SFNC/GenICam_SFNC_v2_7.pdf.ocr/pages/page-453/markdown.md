|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 24.6 ChunkRegionSelector

|  Name | ChunkRegionSelector  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Region0 (if 0 based) Region1 Region2 ... Scan3dExtraction0 (if 0 based) Scan3dExtraction1 Scan3dExtraction2 ...  |

Selects which Region to retrieve data from.

This generally maps to the corresponding RegionSelector feature.

Note that if multiple Regions of interest are supported by the device, the **ChunkRegionSelector** can be added to various features such as ChunkWidth, ChunkHeight... to specify the characteristics of the selected Region. In order to simplify the standard text and feature descriptions, the optional **ChunkRegionSelector** is not propagated to all the features of the SFNC that it can potentially select.

Possible values are:

- **Region0**: Image comes from the Region 0.
- **Region1**: Image comes from the Region 1.
- **Region2**: Image comes from the Region 2.
- ...
- **Scan3dExtraction0**: Image comes from the Scan3dExtraction output Region 0.
- **Scan3dExtraction1**: Image comes from the Scan3dExtraction output Region 1.
- **Scan3dExtraction2**: Image comes from the Scan3dExtraction output Region 2.
- ...