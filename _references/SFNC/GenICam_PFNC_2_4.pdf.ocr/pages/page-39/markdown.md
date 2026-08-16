# Pixel construction rules for the “msb Grouped” style

This packing style is applicable only when each component contains more than 8 bits but no more than 12 bits.

To construct the pixel stream:

1) Take the 8 lsb’s of the first component and put them in the first byte. Reserve the extra msb’s of this component for the last byte(s).
2) Take the 8 lsb’s of the second component and put them in the second byte. Reserve the extra msb’s of this component for the last byte(s).
3) Proceed in this way, taking the 8 lsb’s of the next component and putting it in the next byte until no more components left. For each component, reserve the extra msb’s of this component for the last byte(s). This grouping could stop at the line or image boundary, as explained in section 6.7.
4) Start filling the last byte(s) from its lsb by successively using the extra msb’s from the first component. For monochrome components, add msb padding bits next to the component extra msb’s such that it occupies the indicated number of bits for the monochrome pixel before proceeding with the next component. Continue filling the last byte(s) using the previous rule for each component in turn.
5) Pad the last byte’s msb’s with 0 if needed.

# 6.5 Align tag

Align (“a”) tag can be used to complement the packed and grouped styles. It indicates the total number of bits to align the pixel (if the packing style refers to multi-components) or cluster (if the packing style refers to packing of single-component/monochrome pixels) when there is at least one full byte of padding zeros.

Alignment bits must be set to 0 (they are padding bits). The alignment bytes must be put after any bytes containing component information.

|   | byte 3 | byte 2 | byte 1 | byte 0  |
| --- | --- | --- | --- | --- |
|  ... | 7 . . . . . 0 all 0 (alignment) | 7 . . . . . 0 N | 7 . . . . . 0 M | 7 . . . . . 0 L  |
|  ← | ← | ← | ← | ←  |

Figure 6-17: RGB 8-bit unpacked aligned to 32-bit (RGB8a32)

|   | byte 7 | byte 6 | byte 5 | byte 4 | byte 3 | byte 2 | byte 1 | byte 0  |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
|  ... | Alignment ‘0’ |   | 3^{rd} Mono10 |   | 2^{nd} Mono10 |   | 1^{st} Mono10  |   |
|  ← | ← | ← | ← | ← | ← | ← | ← | ←  |

Figure 6-18 : Using a cluster marker of 3 unpacked Mono10 aligned to 64 bits (Mono10c3a64)