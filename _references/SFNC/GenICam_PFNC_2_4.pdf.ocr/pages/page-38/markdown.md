#### Pixel construction rules for the "lsb Grouped" style

This packing style is applicable only when each component contains more than 8 bits but no more than 12 bits.

To construct the pixel stream:

1) Take the 8 msb's of the first component and put them in the first byte. Reserve the extra lsb's of this component for the last byte(s).
2) Take the 8 msb's of the second component and put them in the second byte. Reserve the extra lsb's of this component for the last byte(s).
3) Proceed in this way, taking the 8 msb's of the next component and putting it in the next byte until no more components left. For each component, reserve the extra lsb's of this component for the last byte(s). This grouping could stop at the line or image boundary, as explained in section 6.7.
4) Start filling the last byte(s) from its lsb by successively using the extra lsb's from the first component. For monochrome components, add msb padding bits next to the component extra lsb's such that it occupies the indicated number of bits for the monochrome pixel before proceeding with the next component. Continue filling the last byte(s) using the previous rule for each component in turn.
5) Pad the last byte's msb's with 0 if needed.

#### 6.4.2 msb Grouped

For msb grouped, the lsb's of the components are extracted and put in sequence starting with the first component in byte 0. The msb's of the components are grouped together in a separate byte that is put last. The principle is the same as lsb grouped. The last byte is filled by grouping components starting from the lsb using the component order, with no empty bit in between. For PFNC-compliant image buffers, msb grouped must be explicitly specified in the pixel format name by appending "msb".

Note that in the following figure, we put byte 0 on the left to help illustrate the concept.

|  byte 0 | byte 1 | byte 2 | byte 3  |
| --- | --- | --- | --- |
|  7 . . . . . . 0L | 7 . . . . . . 0M | 7 . . . . . . 0N | pp 98 98 98N M L ...  |

Figure 6-16 : 3 components of 10-bit with msb grouped into 32-bit pixel (RGB10g32msb)