|  GEN<i>CAM |   | ![img-4.jpeg](img-4.jpeg)emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

|   | 1 | 2 | 3 | 4 | ...  |
| --- | --- | --- | --- | --- | --- |
|  1 | \( LMN_{11} \) | \( LMN_{12} \) | \( LMN_{13} \) | \( LMN_{14} \) | ...  |
|  2 | \( LMN_{21} \) | \( LMN_{22} \) | \( LMN_{23} \) | \( LMN_{24} \) | ...  |
|  3 | \( LMN_{31} \) | \( LMN_{32} \) | \( LMN_{33} \) | \( LMN_{34} \) | ...  |
|  4 | \( LMN_{41} \) | \( LMN_{42} \) | \( LMN_{43} \) | \( LMN_{44} \) | ...  |
|  ... | ... | ... | ... | ... | ...  |

Figure 3-2: LMN444 Pixel Location

#### 3.1.3 LMN422 Location

This format is a 4:2:2 co-sited sub-sampled representation of a 3 component color space. The M and N components are sub-sampled by 2 horizontally: their effective positions are co-sited with alternate L samples, starting in the first column.

Ex: YCbCr422_8

|   | 1 | 2 | 3 | 4 | ...  |
| --- | --- | --- | --- | --- | --- |
|  1 | \( LMN_{11} \) | \( L_{12} \) | \( LMN_{13} \) | \( L_{14} \) | ...  |
|  2 | \( LMN_{21} \) | \( L_{22} \) | \( LMN_{23} \) | \( L_{24} \) | ...  |
|  3 | \( LMN_{31} \) | \( L_{32} \) | \( LMN_{33} \) | \( L_{34} \) | ...  |
|  4 | \( LMN_{41} \) | \( L_{42} \) | \( LMN_{43} \) | \( L_{44} \) | ...  |
|  ... | ... | ... | ... | ... | ...  |

Figure 3-3: LMN422 Pixel Location

When 4:2:2 sub-sampling is used, the components are transmitted using the following order, unless a component order is explicitly stated in the standard referencing the Pixel Format Naming Convention.

\[
\mathrm{L} _ {1 1}, \mathrm{M} _ {1 1}, \mathrm{L} _ {1 2}, \mathrm{N} _ {1 1}, \mathrm{L} _ {1 3}, \mathrm{M} _ {1 3}, \mathrm{L} _ {1 4}, \mathrm{N} _ {1 3} \dots
\]

The above component order is equivalent to FourCC \( ^{1} \) YUY2.