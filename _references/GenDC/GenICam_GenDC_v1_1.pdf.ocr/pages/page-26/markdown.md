|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

### 2.2.8 GenDC Part Header Type Specific Fields

This section describes the Part type specific Header fields. The Part type specific fields are used to describe the type of data included in the Part and to provide the information necessary to interpret it. The sub sections below gives the information about each of the in GenDC Part Header Types defined above.

#### 2.2.8.1 Metadata specific Part Header fields

|  Size  |   |   |
| --- | --- | --- |
|  Padding | PaddingReserved | InfoReserved  |
|  InfoTypeSpecific (Reserved)  |   |   |

Figure 2-7: Part specific Header fields layout for Metadata

Note: For Part Type = Metadata, the Component's Format and Part's Format are set to the Data8 PFNC value.

|  Metadata Part Type specific Header fields  |   |   |
| --- | --- | --- |
|  Width (Bytes) | Offset (Bytes) | Description  |
|  8 | 40 | SizeSize of the 1D Metadata Part (in number of elements according to Part Format).  |
|  2 | 48 | PaddingSize of the padding at the end of the Part (in bytes).  |
|  2 | 50 | PaddingReserved = 0.Reserved for alignment and future use.  |
|  4 | 52 | InfoReserved = 0Reserved for future use.  |
|  8 | 56 | InfoTypeSpecific = 0 or GenICam LayoutId.Reserved for future use (0) or set to the Metadata Layout Id for Part of type GDC_METADATA_GENICAM_CHUNK or M_METADATA_GENICAM_XML (See “GenDC metadata Part layout for GenICam Chunk” below).  |

Table 2-6: Part specific Header fields description for Metadata

##### GenDC Metadata Part:

A GenDC Metadata Part is a Part containing non pixel based information.

There are various types of such data identified by their specific Part Header type.

##### GenDC Metadata Part layout for GenICam Chunk

A GenDC Metadata Part of type GDC_METADATA_GENICAM_CHUNK includes GenICam chunk data that are tagged blocks of data separated in individual information chunks identified by a unique chunk identifier. A data block in chunk format is decoded with the help of a GenICam chunk parser and the corresponding GenICam XML. Each individual