|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

Replacing (8) into (7) leads to:

|  Y' = 0.299 R' + 0.587 G' + 0.114 B'  |
| --- |
|  Cb = -0.16874 R' - 0.33126 G' + 0.5000 B' + 128  |
|  Cr = 0.5000 R' - 0.41869 G' - 0.08131 B' + 128  |
|  with R', G' and B' in the range [0, 255].  |

Equation 2 : Generic full scale R'G'B' to Y'CbCr conversion (8 bits)

For 8-bit data, the valid range of values for each component is:

❖ Y', R', G' and B' in range [0, 255], unsigned (256 levels)
❖ Cb and Cr in range [0, 255], signed shifted by 128, with 128 representing 0 (256 levels)

Values must be truncated to fit in that range.

Note that the above equations are the ones specified by the JFIF specification (JPEG File Interchange Format).

The reverse equations are given by:

|  R' = Y' + 1.40200 (Cr - 128)  |
| --- |
|  G' = Y' - 0.34414 (Cb - 128) - 0.71414 (Cr - 128)  |
|  B' = Y' + 1.77200 (Cb - 128)  |
|  with Y', Cb and Cr in the range [0, 255].  |

Equation 3 : Generic full scale Y'CbCr to R'G'B' conversion (8 bits)

Equivalently, the same set of equations can be used for generic YUV where the range of values for each component must use the full 8-bit. In this case, U = Cb and V = Cr.

### 8.2.2 Y'CbCr601 (8-bit)

ITU-R BT.601 provides a definition of the Y', Cb and Cr based on (1). It defines the following signal range:

Y' in the range [16, 235] (9)

Cb and Cr in the range [16, 240]

Since BT.601 is based on (1), it leads to the same color difference signal indicated in (5) and (6).