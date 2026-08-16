|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

data of an individual Components group. In that case, the information in each Metadata Component applies only to the other Components sharing the same GroupId and the GenICam chunk parser must be called for each of them to retrieve the information about a particular group.

#### GenDC Metadata Part layout for GenICam XML

A GenDC Metadata Part of type GDC_METADATA_GENICAM_XML is formatted block of XML data respecting the GenICam XML format and its schema.

For those GenICam XML Parts, the Part's header Flags field must be used to identify the format and content of XML data of the Part. This can be used for example to specify if the XML is zipped (compressed) or if it is suitable for GenICam Chunk data decoding (see 2.2.4.1 GenDC Part Header Common Fields Description).

Each of the Metadata Components of a Container that has a GDC_METADATA_GENICAM_CHUNK Part can also include a GDC_METADATA_GENICAM_XML Part corresponding to the chunk data of that Component. Or, if suitable to avoid duplication, a common and unique Metadata Component including a single Part of type

GDC_METADATA_GENICAM_XML can be provided for all the Components of a Container containing a GenICam chunk metadata Part. This XML can then be used to decode any of the Container's chunk Metadata Parts. In that case, the Component containing the common GenICam XML must have the GroupId 0xFFFF (which represents "pertains to all GroupIds").

When the content (availability or characteristics of individual XML features) changes in the XML data Part, the XML "LayoutId" contained in the InfoTypeSpecific field of the GenICam Metadata Part header must change too. As long as the XML layout remains the same, the XML LayoutId should stay identical.

Note: It is recommended that the sources that support sending GDC_METADATA_GENICAM_XML Parts in a GenDC Container have this functionality configurable and disabled by default (i.e. SFNC feature ChunkXMLEnable = False).

##### 2.2.8.2 1D Array specific Part type fields

|  Size  |   |   |
| --- | --- | --- |
|  Padding | PaddingReserved | InfoReserved  |
|  InfoTypeSpecific (Reserved)  |   |   |

Figure 2-9: Part specific Header fields layout for 1D Array

|  1D Array (ex: Point Cloud) Part Type specific Header fields  |   |   |
| --- | --- | --- |
|  Width (Bytes) | Offset (Bytes) | Description  |
|  8 | 40 | SizeSize of the 1D Part (in number of elements according to Part Format).  |
|  2 | 48 | PaddingSize of the padding at the end of the Part (in bytes).  |