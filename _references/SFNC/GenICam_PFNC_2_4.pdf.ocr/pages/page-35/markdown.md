|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

The following example shows how a 10-bit monochrome data can be packed from its lsb, again with no padding bit.

|   | byte 4 | byte 3 | byte 2 | byte 1 | byte 0  |
| --- | --- | --- | --- | --- | --- |
|  ... | 9 . . . . . . 2L4 | 1 0 9 . . . . 4L4 L3 | 3 . . 0 9 . . 6L3 L2 | 5 . . . . 0 9 8L2 L1 | 7 . . . . . . 0L1  |

Figure 6-9 : 10-bit monochrome pixel lsb packed (Mono10p)

#### Pixel construction rules for the "lsb Packed" style

The total number of bits after packing is either:

1) Indicated by the number following the "p" tag when present
2) Deduced by putting as many components such that no padding bit is required.

To construct the pixel stream:

1) Take the first component and put it in the lsb's of the first byte, with bit 0 holding the lsb of the component. Extra bits of this component continue in the lsb's of the next byte.
2) Then take the following component and append it to the first one, again starting from the lsb of the component.
3) Proceed in this way, appending the next component from its lsb, until no more components left.
4) Pad the last byte's most significant bits with 0 if needed (i.e. to meet the total number of bits indicated after the "p" tag). This padding must consider the line or image boundary, as explained in section 6.7.

#### 6.3.2 msb Packed

For msb packed, the data is filled msb first in the lowest address byte (byte 0), starting with the first component. For PFNC-compliant image buffers, msb packed must be explicitly specified in the pixel format name by appending “msb” after the ‘p’ (i.e. “pmsb”).

Note that in the following figure, we put byte 0 on the left to help illustrate the concept. The data is filled msb first in the lowest address byte (byte 0) starting with the first component and continue in the msb of byte 1 (and so on). Padding bits, if any, would thus be the lsb's of the last byte after putting all the components. Padding bits are necessary when the "p" packing tag is followed by a number indicating to how many bits we need to align the pixel.