|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- Ultraviolet: The image data is the ultraviolet component.
- Range: The image data is the range component (distance or depth).
- Reflectance: The image data is the reflected intensity component (acquired with the Range).
- Disparity: The image data is the disparity component.
- Confidence: The image data is the confidence map component.
- Scatter: The image data is the scatter component.
- Multispectral: The image data is the multispectral component.

### 24.10 ChunkComponentID

|  Name | ChunkComponentID  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Intensity Infrared Ultraviolet Range Reflectance Disparity Confidence Scatter Multispectral Device-specific  |

Returns the Identifier of the selected Component. This can be used to identify the image component type of a multi-component payload.

For example, the Intensity and Scatter Components of a multi-component data buffer may be of the same pixel format and impossible to distinguish without this information.

This generally maps to the corresponding ComponentSelector feature.

Possible values are: