|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

### 2.2.7 GenDC Part Header Types

This section lists the defined Part Header types.

Value = 0x4000-0x4FFF, 4 bits are used for the general category; the following low 8 bits are used for the sub type of a Part. It is possible to specify a custom category (0x4Fxx), which gives 256 possible custom sub types. Besides this, each standardized category supports up to 16 custom sub types (0x4xFx).

|  Type | Value | Description  |
| --- | --- | --- |
|  GDC_GENERIC_PART_METADATA | 0x40xx | Generic Metadata Part type.  |
|  GDC_METADATA_GENICAM_CHUNK | 0x4000 | GenICam Chunk Metadata. Binary chunk metadata formatted as specified in the “GenDC metadata Part layout for GenICam Chunk” section below.Note:- Only a Component with a TypeId of type Metadata can contain a Part of type GDC_METADATA ...  |
|  GDC_METADATA_GENICAM_XML | 0x4001 | GenICam Metadata XML. GenICam XML formatted as specified in the GenICam standard.Note:- If the Part XML data is zipped (compressed) or if it is to be used for GenICam chunk data decoding, the Part header Flags field bits “Zip” and “Chunk” must be set accordingly.- Only a Component with a TypeId of type Metadata can contain a Part of type GDC_METADATA ...  |
|  GDC_METADATA_CUSTOM(x) | 0x40Fx | The Metadata Part data is custom and does not correspond to a known Type. It is uniquely identified by the lower 4 bits of this field.  |
|  GDC_GENERIC_PART_1D | 0x41xx | Generic 1D Part type.  |
|  GDC_1D | 0x4100 | 1D array (such as 3D Point Cloud).  |
|  GDC_1D_CUSTOM(x) | 0x41Fx | The 1D Part data is custom and does not correspond to a known Type. It is uniquely identified by the lower 4 bits of this field.  |
|  GDC_GENERIC_PART_2D | 0x42xx | Generic 2D Part type.  |
|  GDC_2D | 0x4200 | Rectangular uncompressed image (monochrome or none planar color).  |
|  GDC_2D_JPEG | 0x4201 | JPEG compressed Image.  |
|  GDC_2D_JPEG2000 | 0x4202 | JPEG 2000 compressed Image.  |