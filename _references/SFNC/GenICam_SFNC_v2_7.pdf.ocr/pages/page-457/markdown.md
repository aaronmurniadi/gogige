|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- Intensity: The image data is the intensity component (visible).
- Infrared: The image data is the infrared component.
- Ultraviolet: The image data is the ultraviolet component.
- Range: The image data is the range component (distance or depth).
- Reflectance: The image data is the reflected intensity component (acquired with the Range).
- Disparity: The image data is the disparity component.
- Confidence: The image data is the confidence map component.
- Scatter: The image data is the scatter component.
- Multispectral: The image data is the multispectral component.

### 24.11 ChunkComponentIDValue

|  Name | ChunkComponentIDValue[ChunkComponentSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns a unique Identifier value that corresponds to the selected chunk Component.

This value generally maps to corresponding ComponentIDValue feature value.

### 24.12 ChunkGroupSelector

|  Name | ChunkGroupSelector  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |