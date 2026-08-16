|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

#### 2.2.8.4 H.264 compressed image specific Part Header fields

|  SizeX |   |   |   |   |   | SizeY  |   |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  PaddingX |   | PaddingY |   |   |   | InfoReserved  |   |
|  Reserved | ProfileIDC | CS | PM | RF | LevelIDC | SpropInterleavingDepth | SpropMaxDonDiff  |
|  SpropDeintBufReq |   |   |   |   |   | SpropInitBufTime  |   |

Figure 2-11: Part specific Header fields layout for H.264 compressed image

|  2D H.264 Compressed Part Type specific Header fields  |   |   |
| --- | --- | --- |
|  Width (Bytes) | Offset (Bytes) | Description  |
|  4 | 40 | SizeX X size of the 2D Part (in pixels).  |
|  4 | 44 | SizeY Y size of the 2D Part (in pixels).  |
|  2 | 48 | PaddingX Size of the X padding at the end of each line (in bytes).  |
|  2 | 50 | PaddingY Size of the Y padding at the end of the Part (in bytes).  |
|  4 | 52 | InfoReserved = 0 Reserved for future use.  |
|  1 | 56 | Reserved = 0 Reserved for future use.  |
|  1 | 57 | ProfileIDC Profile IDC sequence parameter set data attribute as defined by H.264.  |