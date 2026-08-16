|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

|  2 | 50 | PaddingReserved = 0. Reserved for alignment and future use.  |
| --- | --- | --- |
|  4 | 52 | InfoReserved = 0 Reserved for future use.  |
|  8 | 56 | InfoTypeSpecific = 0. Reserved for future use (0).  |

Table 2-7: Part specific Header fields description for 1D Array or Image

### 2.2.8.3 2D uncompressed, JPEG or JPEG2000 compressed image specific Part type fields

|  SizeX |   | SizeY  |
| --- | --- | --- |
|  PaddingX | PaddingY | InfoReserved  |

Figure 2-10: Part specific Header fields layout for 2D uncompressed, JPEG or JPEG2000 compressed image

|  Width (Bytes) | Offset (Bytes) | Description  |
| --- | --- | --- |
|  4 | 40 | SizeX X size of the 2D Part (in pixels).  |
|  4 | 44 | SizeY Y size of the 2D Part (in pixels).  |
|  2 | 48 | PaddingX Size of the X padding at the end of each line (in bytes).  |
|  2 | 50 | PaddingY Size of the Y padding at the end of the Part (in bytes). Padding Y can be used to align the following Part to specific hardware constraints, e.g. processor specific alignment constraints  |
|  4 | 52 | InfoReserved = 0 Reserved for future use.  |

Table 2-8: Part specific Header fields description for 2D uncompressed, JPEG or JPEG2000 compressed image