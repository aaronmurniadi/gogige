**Pixel construction rules for the “Align” style**

This packing style is applicable only when at least one full byte contains padding bits and alignment must be on an 8-bit boundary.

To construct the pixel stream:

1) Use the unpacked, packed or grouped style specified by the pixel format using the pixel construction rules above.
2) Pad the most significant bits with as many padding bits required to align the data to the number of bits specified by the ‘a’ tag.