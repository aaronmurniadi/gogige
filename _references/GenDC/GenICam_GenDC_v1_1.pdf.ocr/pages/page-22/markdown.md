|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

|  4 | 8 | **Format** = Data format of the Part. The value is specified using the standard PFNC's Pixel Format Values list. See the GenICam Pixel Format Naming Convention's "Pixel Format Values" document. In general, the Part's Format is identical to its encapsulating Component's Format. For a Component that has a planar format, the Parts of the planar buffer must be ordered according to their encapsulating Component's Format. The Part's Format describes the format of each individual Part/plane of the encapsulating Component (e.g. For a PFNC_RGB8_planar Component, the individual Parts format will be PFNC_R8, PFNC_G8 and PFNC_B8). For Parts that are not pixel based like metadata, the PFNC_Data8 format must be used.  |
| --- | --- | --- |
|  2 | 12 | **Reserved** = Reserved for future use (set to 0).  |
|  2 | 14 | **FlowId** Unique identifier of the data Flow used to transport and store the Part's data. Start at 0 and incrementing. See chapter 3 for more information.  |
|  8 | 16 | **FlowOffset** Offset in bytes of the Part's data in the data Flow used to transport and store the Part's data relative to the Flows' base address. The FlowOffset of a Part is equal to the DataOffset of the Part minus the base offset of the Flow in which it is located. Note that the Flow base addresses are not part of the Container Descriptor as the Container must be independent of the actual storage location. This is the same as for the Containers' Descriptor base address which is identical to the base address of Flow 0. If FlowId=0 it is therefore the same as DataOffset. For other values of FlowId, DataOffset gives the Flows' start address for linear addressing if FlowOffset=0. This can be used to reconstruct a linear Container e.g. for storage.  |
|  8 | 24 | **DataSize** = Size of the Part's data in bytes. Typically the maximum possible data size for a Part. In the final Descriptor the Part's real and valid data size (line scan variable SizeY, compressed image, ...).  |
|  8 | 32 | **DataOffset** Offset in bytes of the start of the Part's data section relative to the start of the Container's Descriptor for linear addressing.  |

Table 2-3: GenDC Component Header Part common fields description

### 2.2.6.2 GenDC Part Header Type Specific Fields Description