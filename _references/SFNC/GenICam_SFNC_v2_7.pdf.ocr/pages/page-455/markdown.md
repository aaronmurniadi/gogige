|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Interface | Integer  |
| --- | --- |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns the unique integer Identifier value of the Region that the image comes from.

This generally maps to the corresponding RegionIDValue feature.

### 24.9 ChunkComponentSelector

|  Name | ChunkComponentSelector  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Intensity Infrared Ultraviolet Range Reflectance Disparity Confidence Scatter Multispectral Device-specific  |

Selects the Component from which to retrieve data from.

This generally maps to the corresponding ComponentSelector feature.

Possible values are:

- Intensity: The image data is the intensity component (visible).
- Infrared: The image data is the infrared component.