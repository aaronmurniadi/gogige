|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

#### 3.1.4 LMN411 Location

This format is a 4:1:1 co-sited sub-sampled representation of a 3 component color space. The M and N components are sub-sampled by 4 horizontally and are thus associated to 4 consecutive columns. Their position is co-sited starting with the first L sample.

Ex: YCbCr411_8

|   | 1 | 2 | 3 | 4 | 5 | ...  |
| --- | --- | --- | --- | --- | --- | --- |
|  1 | \( LMN_{11} \) | \( L_{12} \) | \( L_{13} \) | \( L_{14} \) | \( LMN_{15} \) | ...  |
|  2 | \( LMN_{21} \) | \( L_{22} \) | \( L_{23} \) | \( L_{24} \) | \( LMN_{25} \) | ...  |
|  3 | \( LMN_{31} \) | \( L_{32} \) | \( L_{33} \) | \( L_{34} \) | \( LMN_{35} \) | ...  |
|  4 | \( LMN_{41} \) | \( L_{42} \) | \( L_{43} \) | \( L_{44} \) | \( LMN_{45} \) | ...  |
|  ... | ... | ... | ... | ... | ... | ...  |

Figure 3-4: LMN411 Pixel Location

When 4:1:1 sub-sampling is used, the components are transmitted using the following order, unless a component order is explicitly stated in the standard referencing the Pixel Format Naming Convention.

\[
\mathrm{L} _ {1 1}, \mathrm{L} _ {1 2}, \mathrm{M} _ {1 1}, \mathrm{L} _ {1 3}, \mathrm{L} _ {1 4}, \mathrm{N} _ {1 1}, \mathrm{L} _ {1 5}, \mathrm{L} _ {1 6}, \mathrm{M} _ {1 5}, \mathrm{L} _ {1 7}, \mathrm{L} _ {1 8}, \mathrm{N} _ {1 5} \dots
\]

#### 3.1.5 LMNO4444 Location

This format is typically used for any 4 component color space, such as aRGB (where 'a' represents alpha compositing). No sub-sampling is performed.

Ex: aRGB8

|   | 1 | 2 | 3 | 4 | ...  |
| --- | --- | --- | --- | --- | --- |
|  1 | \( LMNO_{11} \) | \( LMNO_{12} \) | \( LMNO_{13} \) | \( LMNO_{14} \) | ...  |
|  2 | \( LMNO_{21} \) | \( LMNO_{22} \) | \( LMNO_{23} \) | \( LMNO_{24} \) | ...  |
|  3 | \( LMNO_{31} \) | \( LMNO_{32} \) | \( LMNO_{33} \) | \( LMNO_{34} \) | ...  |
|  4 | \( LMNO_{41} \) | \( LMNO_{42} \) | \( LMNO_{43} \) | \( LMNO_{44} \) | ...  |
|  ... | ... | ... | ... | ... | ...  |

Figure 3-5: LMNO4444 Pixel Location