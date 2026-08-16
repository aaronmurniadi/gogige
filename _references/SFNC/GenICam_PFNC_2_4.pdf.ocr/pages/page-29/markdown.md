|  **GEN<i>CAM** |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

## 4 Number of bits for each component

This field provides the number of bits for each component. Typical values are

- ○ 1, 2, 4, 5, 6, 8, 10, 12, 14 and 16 for integer data types;
- ○ 32 and 64 for floating point data types.

Use one number if all components of the pixel have same number of bits (ex: Mono8), otherwise one must successively list one number for each component with no space in-between using the same number of digits for all components (including a leading zero when necessary)

Ex: RGB565 = 5-bit R + 6-bit G + 5-bit B
YCbCr160808 = 16-bit Y' + 8-bit Cb + 8-bit Cr

From the above, one can deduce the number of bits occupied by the pixel (not including padding bits). If a single value is listed, then the number of bits is equal to number of components multiplied by the number of bits:

Ex: RGB8 = 3 components of 8-bit = 24 bits
Coord3D_AC16 = 2 components of 16-bit = 32 bits

When the components don't use the same number of bits, then it is the concatenation of them:

Ex: RGB565 = 5-bit for red + 6-bit for G + 5-bit for B = 16 bits for each pixel

The "Packing Style" section introduces padding bits that increases to overall size of the pixel. In those situations where padding bits are used, the packing style might include a number representing the number of bits used by the pixel, including zero-padding.