|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

Each Part can have multiple 64-bit TypeSpecific fields (shown in blue). For all the Part types defined in this document, these fields are described in the section 2.2.7. In general, if applicable, those TypeSpecific fields should follow a common layout. This common layout illustrated in Figure 2-6, starts with the Parts' dimensions (8 bytes, e.g. for 2D images SizeX and SizeY each 4 bytes), followed by the Parts' padding (4 bytes, e.g. for 2D images PaddingX and PaddingY each 2 bytes) and a reserved field (4 bytes) for future additions. Other optional Part Type specific fields (8 bytes each) can also be present.

|  Dimension  |   |
| --- | --- |
|  Padding | InfoReserved  |
|  TypeSpecific 3 (optional)  |   |
|  ...  |   |
|  TypeSpecific n (optional)  |   |

Figure 2-6: GenDC Part TypeSpecific fields general layout