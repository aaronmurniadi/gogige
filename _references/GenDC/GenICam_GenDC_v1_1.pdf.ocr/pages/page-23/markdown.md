|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

This table describes the Part's TypeSpecific fields as presented in Figure 2-6: GenDC Part TypeSpecific fields general layout. If applicable, those fields should use the following general definition. This permits consistent and easier to interpret Parts headers. In chapter 2.2.8 GenDC Part Header Type Specific Fields, you can find Part specific definitions for those generic fields.

|  Width (Bytes) | Offset (Bytes) | Description  |
| --- | --- | --- |
|  8 | 40 | **Dimension** Typically the size of the Part (e.g. For 2D images SizeX and SizeY. For 1D array SizeX only).  |
|  4 | 48 | **Padding** Typically the padding used by the Part (e.g. for 2D images PaddingX and PaddingY. For 1D array PaddingX only).  |
|  4 | 52 | **InfoReserved** For future generic Part info use. Set to 0.  |
|  8xN | 56 | **TypeSpecific fields 3 to N (optional)** Optional fields to be used to further describe the Part if necessary. Note that depending on the Type the above 3 fields might not exist. In that case, the TypeSpecific fields will start from Offset 40 instead. This might especially be true for custom formats. It is however highly recommended to use the general layout starting with Dimension and Padding even for custom Parts since generic software can make use of it. For example even if a specific Part Header type is unknown, but it is a 2D image, a generic software can present the data to the user as raw data but with correct SizeX and SizeY.  |

Table 2-4: GenDC Component Header Part type specific fields description