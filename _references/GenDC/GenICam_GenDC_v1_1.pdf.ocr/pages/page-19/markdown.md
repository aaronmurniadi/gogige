|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

|  8 | 32 | **TypeId** Component type identifier. This number uniquely identifies what type the Component data represents and is equal to one of the SFNC's ComponentIdValue feature predefined values.  |
| --- | --- | --- |
|  4 | 40 | **Format** Format of the whole Component (including all its Parts). The value is specified using the standard PFNC's Pixel Format Values list. See the GenICam Pixel Format Naming Convention's "Pixel Format Values" document. If the Component has only a single Part the Component's Format is the same as the Part's Format. For planar formats it is the encapsulating format (e.g. RGB8_planar). For the Components that are not pixel based like metadata, the PFNC_Data8 format must be used.  |
|  2 | 44 | **Reserved2** = Reserved for future use (set to 0).  |
|  2 | 46 | **PartCount** Number of Parts in the Component. Note that planar Components must include one separate Part per data plane.  |
|  PartCount x 8 | 48 | **PartOffset[]** Array of the offsets in bytes of the start of each of the Part Headers relative to the start of the Containers's Header. The size of the array is **PartCount** x 8 bytes.  |

Table 2-2: GenDC Component Header fields description

### 2.2.5 GenDC Part Header Layout

The following figure shows the layout of the Part Headers. In white common fields to all Parts (see section 2.2.6.1) and in blue the fields which are specific to each particular Part Type (see section 2.2.8).

|  HeaderType | Flags | HeaderSize  |   |
| --- | --- | --- | --- |
|  Format |   | Reserved | FlowId  |
|  FlowOffset  |   |   |   |
|  DataSize  |   |   |   |
|  DataOffset  |   |   |   |
|  TypeSpecific 1 (optional)  |   |   |   |
|  ...  |   |   |   |
|  TypeSpecific n (optional)  |   |   |   |

Figure 2-5: GenDC Part Header Layout