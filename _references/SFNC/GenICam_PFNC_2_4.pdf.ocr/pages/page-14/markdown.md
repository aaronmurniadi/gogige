|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

## 3 Components and Location

The Components field provide the list of component constituents available in the pixel format. Location offers additional information regarding the positioning of the components in the image. The combination of the two gives a good idea of the pixel format as seen by the user.

### 3.1 Pixel Location in Image

This section lists the various components positioning within the image. It is especially helpful when sub-sampling of certain components is used. This information is required to determine the “Components and Location” field of the pixel name.

In the following diagrams, for a given pixel, the first index represents the row number; the second index represents the column number.

The figures of this section uses generic pixel component format where 'L' represents the first component, 'M' the second, 'N' the third and 'O' the fourth (if necessary). To help clarify some of them, you can think about LMN = RGB (where R = L, G = M and B = N) or LMN = Y'CbCr (where Y' = L, Cb = M and Cr = N). Same hold true for Bayer patterns (where R = L, G = M and B = N).

#### 3.1.1 Mono Location

This format is used for single component images where typically L is the luma (Y'). This could also be used for planar transfer where each component of the pixel is separated onto a different stream.

Ex: Mono8

|   | 1 | 2 | 3 | 4 | ...  |
| --- | --- | --- | --- | --- | --- |
|  1 | \( L_{11} \) | \( L_{12} \) | \( L_{13} \) | \( L_{14} \) | ...  |
|  2 | \( L_{21} \) | \( L_{22} \) | \( L_{23} \) | \( L_{24} \) | ...  |
|  3 | \( L_{31} \) | \( L_{32} \) | \( L_{33} \) | \( L_{34} \) | ...  |
|  4 | \( L_{41} \) | \( L_{42} \) | \( L_{43} \) | \( L_{44} \) | ...  |
|  ... | ... | ... | ... | ... | ...  |

Figure 3-1: Mono Pixel Location

#### 3.1.2 LMN444 Location

This format is typically used for any 3 component color space, such as RGB and Y'CbCr. No sub-sampling is performed.

Ex: RGB8