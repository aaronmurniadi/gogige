### 3.2.1 GenDC Flow mapping table

The Transport Layer gets the Flow mapping table that provides the Flow information about the number of Flows and the size of each Flow in a Transport Layer specific way, for example using the device XML or bootstrap registers. The Flow information is static after TLParamsLocked has been set.

|  HeaderType |   | Flags | HeaderSize  |
| --- | --- | --- | --- |
|  VersionMajor | VersionMinor | Reserved | FlowCount  |
|  FlowSize[FlowCount] ...  |   |   |   |

Figure 3-3: GenDC Flow mapping table layout

|  Width (Bytes) | Offset (Bytes) | Description  |   |   |
| --- | --- | --- | --- | --- |
|  2 | 0 | **HeaderType** = Unique Header format identifier (Flow Mapping Table Header) (GDC_FLOW_MAPPING_HEADER = 0x7000).  |   |   |
|  2 | 2 | **Flags** = Flags specifying the characteristics and format of the Table.  |   |   |
|   |   |  Width (bits) | Bit offset (lsb << x) | Description  |
|   |   |  16 | 0 | Reserved (set to 0)  |
|  4 | 4 | **HeaderSize** = Size of the Flow Mapping Table Size of the Flow Mapping Table Header in bytes including the variable sized FlowSize array (i.e.16+ FlowCount x 8)  |   |   |
|  1 | 8 | **VersionMajor** = Major Version of this table. Must be set to 1 for this GenDC specification.  |   |   |
|  1 | 9 | **VersionMinor** = Minor Version of this table. Must be set to 0 for this GenDC specification.  |   |   |
|  2 | 10 | **Reserved** = Reserved for future use (set to 0).  |   |   |
|  4 | 12 | **FlowCount** = Number of entries (=Flows) in the table  |   |   |
|  FlowCount x 8 | 16 | **FlowSize**[] Array of the size in bytes of each Flow. The size of the array is **FlowCount** x 8 bytes.  |   |   |

Table 3-1: GenDC Flow mapping table description

The GenDC Container is transmitted using linear addressing relative to the Container start. The Transport Layer can keep this and assign the Flow base addresses also in a linear way or it can use any other Flow base addresses for