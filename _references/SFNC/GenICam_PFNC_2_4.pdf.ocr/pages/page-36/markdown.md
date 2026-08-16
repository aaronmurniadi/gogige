|  byte 0 | byte 1 | byte 2 | byte 3  |
| --- | --- | --- | --- |
|  9 . . . . . . 2L | 1 0 9 . . . . 4L M | 3 . . 0 9 . . 6M N | 5 . . . . 0 p pN  |

Figure 6-10 :3 components in 10-bit msb packed into 32-bit pixel (RGB10p32msb)

|  byte 0 | byte 1 | byte 2 | byte 3 | byte 4  |
| --- | --- | --- | --- | --- |
|  9 . . . . . . 2L1 | 1 0 9 . . . . 4L1 L2 | 3 . . 0 9 . . 6L2 L3 | 5 . . . . 0 9 8L3 L4 | 7 . . . . . . 0L4  |

Figure 6-11 : 10-bit monochrome pixel msb packed (Mono10pmsb)

#### Pixel construction rules for the "msb Packed" style

The total number of bits after packing is either:

1) Indicated by the number following the "p" tag when present; or
2) Deduced by putting as many components such that no padding bit is required.

To construct the pixel stream:

1) Take the first component and put it in the msb's of the first byte, with bit 7 holding the msb of the component. Extra bits of this component continue in the msb's of the next byte.
2) Then take the following component and append it to the first one, again starting from the msb of the component.
3) Proceed in this way, appending the next component from its msb, until no more components left.
4) Pad the last byte's least significant bits with 0 if needed (i.e. to meet the total number of bits indicated after the "p" tag). This padding must consider the line or image boundary, as explained in section 6.7.

### 6.4 Grouped tag

Grouped (“g”) is a different packing style created by regrouping extra lsb’s or msb’s of components (or from successive pixels) in a separate byte(s). The format indicates the number of bits the data occupies when it is different than the nominal bits per pixel for the given component (i.e. including the padding bits). ex: g12 when grouped into 12 bits. This is followed by an optional grouping order indicating if the byte containing the extra data is the lsb’s or msb’s. Empty bit must be padded with 0. The first component is put in byte 0, second component in byte 1 and so on.

When grouped style is used, the byte holding the grouped data shall be put as the last byte(s).