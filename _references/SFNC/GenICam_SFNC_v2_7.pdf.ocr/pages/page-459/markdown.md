|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Level | Optional  |
| --- | --- |
|  Interface | Integer  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns a unique Identifier value that corresponds to the Group of Components of the selected chunk Component.

This value generally maps to corresponding GroupIDValue feature value.

### 24.15 ChunkImageComponent (Deprecated)

|  Name | ChunkImageComponent  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | Intensity Infrared Ultraviolet Range Reflectance Disparity Scatter Confidence Device-specific  |

This feature is deprecated (See ChunkComponentID). It was representing the component of the payload image.

To help backward compatibility, this feature can be included as Invisible in the device's XML.