|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- BayerGB12Packed: Bayer GB 12 bit packed (GigE Vision Specific).
- BayerBG12Packed: Bayer BG 12 bit packed (GigE Vision Specific).
- RGB10V1Packed: RGB 10 bit packed (GigE Vision Specific).
- RGB12V1Packed: RGB 12 bit packed (GigE Vision Specific).
• ...

Note that only a subset of the possible pixel formats is listed here.

See the PixelFormat feature for more details.

24.23 ChunkPixelDynamicRangeMin

|  Name | ChunkPixelDynamicRangeMin  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns the minimum value of dynamic range of the image included in the payload.

24.24 ChunkPixelDynamicRangeMax

|  Name | ChunkPixelDynamicRangeMax  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns the maximum value of dynamic range of the image included in the payload.