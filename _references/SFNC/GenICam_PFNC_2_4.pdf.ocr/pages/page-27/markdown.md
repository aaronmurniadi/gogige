|  Component designation | Positioning in Image | Description  |
| --- | --- | --- |
|  “SCF1WGWB” | SCF1_LNLM location | Sparse Color Filter #1, White-Green-White-Blue pattern  |
|  “SCF1WRWG” | SCF1_LOLN location | Sparse Color Filter #1, White-Red-White-Green pattern  |
|  “SCF1WGWR” | SCF1_LNLO location | Sparse Color Filter #1, White-Green-White-Red pattern  |
|  “CIELAB” | LMN444 location | CIE 1976 L*a*b* color space  |
|  “CIEXYZ” | LMN444 location | CIE 1931 XYZ color space  |
|  “HSI” | LMN444 location | Hue, Saturation, Intensity  |
|  “HSV” | LMN444 location | Hue, Saturation, Value  |
|  “Coord3D_ABC” | LMN444 location | Used for 3D imaging data. Coordinates of the 3D pixel. The depth/range is always represented by coordinate C. *Note: the coordinate system (meaning and unit of individual coordinates) is defined through other means by each device as specified in the GenICam SFNC standard.*  |
|  “Coord3D_A” | Mono location | Used for 3D imaging data. Coordinate A only.  |
|  “Coord3D_B” | Mono location | Used for 3D imaging data. Coordinate B only.  |
|  “Coord3D_C” | Mono location | Used for 3D imaging data. Coordinate C only (the coordinate expressing depth/range). *Note: if the C coordinate is transferred alone, the other two coordinates are implicit as specified in the GenICam SFNC standard.*  |
|  “Coord3D_AC” | LM44 location | Used for 3D imaging data. Coordinates A and C of the 3D pixel. *Note: intended for line scan 3D devices. The second coordinate is implicit as specified in the GenICam SFNC standard.*  |
|  “Confidence” | Mono location | Confidence of the pixel value. Expresses the level of validity of the given pixel value.  |

Note 1: The full scale R, G or B (256 values) can be replaced by their scaled down version r, g or b (235 values) when necessary.

Unless specified otherwise, the order in which the components are listed is the order they will appear on the wire or in memory. The first component appears in the first byte(s) and so on.