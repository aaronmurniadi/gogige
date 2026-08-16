|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

### 2.2.6 GenDC Part Header Description

[ R-002] A GenDC compliant product must use the Part header types as defined by this specification.

#### 2.2.6.1 GenDC Part Header Common Fields Description

This table describes the Part's common fields as presented in Figure 2-5: GenDC Part Header Layout.

|  Width (Bytes) | Offset (Bytes) | Description  |   |   |
| --- | --- | --- | --- | --- |
|  2 | 0 | HeaderType = Part Type Header format identifier (GDC_PART_HEADER = 0x4xxx).A GenDC Component must contain at least one Part Header.See section 2.2.7 for the table of defined Part types.  |   |   |
|  2 | 2 | Flags = Part specific flagsFlags specifying the characteristics and format of the Part.For Part type GDC_METADATA_GENICAM_XML:  |   |   |
|   |   |  Width (bits) | Bit offset (lsb << x) | Description  |
|   |   |  8 | 0 | Reserved (set to 0).  |
|   |   |  1 | 8 | ZipThe Part XML Data is zipped (compressed) according the Zip format used for GenICam XML.  |
|   |   |  1 | 9 | ChunkThe Part XML data can be used for GenICam Chunk data Part decoding.  |
|   |   |  6 | 10 | Reserved = Reserved (set to 0).  |
|   |   |  For all the other Part types:  |   |   |
|   |   |  Width (bits) | Bit offset (lsb << x) | Description  |
|   |   |  16 | 0 | Reserved (set to 0).  |
|  4 | 4 | HeaderSize = Size of the Part type specific Header.Size in bytes of the Part Header including Part Type specific fields.  |   |   |