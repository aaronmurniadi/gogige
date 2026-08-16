|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

### 6.6 Packing Style Summary

Table 6-1 : Packing Style Summary

|  Unpacked | Most significant bits of each component are padded with zeros to be 8-bit aligned.  |
| --- | --- |
|  Unpacked msb | Least significant bits of each component are padded with zeros to be 8-bit aligned.  |
|  “c” | Cluster of ‘x’ single-component/monochrome pixels to be regrouped and consider as one multi-component pixel. This marker must appear before the other tags. Used in conjunction with another packing style (packed, grouped or aligned).  |
|  “p” | lsb packing on ‘x’ bits. Components are packed with no bit spacing, lsb’s in the first byte. Byte 0’s lsb contains the lsb of the first component. Next component is appended going from lsb to msb. Padding bits, if any, occupies the msb’s. If packing style does not explicitly include lsb or msb, then lsb is assumed. ‘x’ indicates the number of bits consumed by the pixel, including the padding 0. ‘x’ is not necessary when there is no padding bit in the resulting pixel data word.  |
|  “pmsb” | msb packing on ‘x’ bits Components are packed with no bit spacing, msb’s in the first byte. Byte 0’s msb contains the msb of the first component. Next component is appended going from msb to lsb. Padding bits, if any, occupies the lsb’s. ‘x’ indicates the number of bits consumed by the pixel, including the padding 0. ‘x’ is not necessary when there is no padding bit in the resulting pixel data word.  |
|  “g” | lsb grouping style on ‘x’ bits Least significant bits of the components are grouped together in a separate byte(s). If packing style does not explicitly include lsb or msb, then lsb is assumed. Byte 0 contains the first component. ‘x’ indicates the number of bits consumed by the pixel, including the padding 0. ‘x’ is not necessary when there is no padding bit in the resulting pixel data word.  |
|  “gmsb” | msb grouping style on ‘x’ bits Most significant bits of the components are grouped together in a separate byte(s). Byte 0 contains the first component. ‘x’ indicates the number of bits consumed by the pixel, including the padding 0. ‘x’ is not necessary when there is no padding bit in the resulting pixel data word.  |
|  “a” | The pixel (or group of pixels) is aligned to the given bit-boundary. This can be used to complement unpacked, packed or grouped packing style. This tag must appear after any other packing style tags. ‘x’ indicates the total number of bits to use for this grouping. Unused bit are set to 0. For multiple single-component/monochrome pixels that must be aligned together, it is mandatory to use the cluster marker (‘c’).  |