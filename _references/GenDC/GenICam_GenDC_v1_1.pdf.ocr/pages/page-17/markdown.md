|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

### 2.2.3 GenDC Component Header Layout

|  HeaderType | Flags | HeaderSize  |   |
| --- | --- | --- | --- |
|  Reserved | GroupId | SourceId | RegionId  |
|  RegionOffsetX |   | RegionOffsetY  |   |
|  Timestamp  |   |   |   |
|  TypeId  |   |   |   |
|  Format |   | Reserved2 | PartCount  |
|  PartOffset[PartCount]  |   |   |   |

Figure 2-4: GenDC Component Layout

### 2.2.4 GenDC Component Header Description

#### 2.2.4.1 GenDC Component Header Common Fields Description

|  Width (Bytes) | Offset (Bytes) | Description  |   |   |
| --- | --- | --- | --- | --- |
|  2 | 0 | **HeaderType** = Unique Header format identifier (Component Header) (GDC_COMPONENT_HEADER = 0x2000). A GenDC Container must always contain at least one Component Header.  |   |   |
|  2 | 2 | **Flags** Flags specifying the characteristics and format of the Component.  |   |   |
|   |   |  Width (bits) | Bit offset (lsb << x) | Description  |
|   |   |  1 | 0 | **Invalid** The Component is invalid and must not be used. If this flag is set the **ComponentInvalid** flag of the Container Header must be also set.  |
|   |   |  15 | 1 | Reserved (set to 0)  |
|  4 | 4 | **HeaderSize** = Size of the Component Header Size of the Component Header in bytes including the variable sized PartOffset array. Note the Component's Part Headers are **not** included.  |   |   |
|  2 | 8 | **Reserved** = Reserved for future use (set to 0).  |   |   |