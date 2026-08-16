|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

### 8.2 Y'CbCr Conversions

Many variants of Y'CbCr exist. Most of them are derived from 2 ITU-R specifications:

1. BT.601 (Studio encoding parameters of digital television for standard 4:3 and wide screen 16:9 aspect ratios) used for standard television
2. BT.709 (Parameter values for the HDTV standards for production and international programme exchange) used for high definition television

Conversions are typically performed starting from the RGB format. Some of the complexity comes from the range of values permitted for each color component. The computer world typically uses full scale value (i.e. 256 values in 8-bit), while BT.601 and BT.709 use a scaled down range spanning 220 values for Y' and 225 values for Cb and Cr.

BT.601 defines the following basic equation for luma in the analog domain:

\[
\mathrm{E} ^ {\prime} \mathrm{Y} = 0. 2 9 9 \mathrm{E} ^ {\prime} \mathrm{R} + 0. 5 8 7 \mathrm{E} ^ {\prime} \mathrm{G} + 0. 1 1 4 \mathrm{E} ^ {\prime} \mathrm{B} \tag {1}
\]

BT.709 defines the following basic equation for luma in the analog domain:

\[
\mathrm{E} ^ {\prime} \mathrm{Y} = 0. 2 1 2 6 \mathrm{E} ^ {\prime} \mathrm{R} + 0. 7 1 5 2 \mathrm{E} ^ {\prime} \mathrm{G} + 0. 0 7 2 2 \mathrm{E} ^ {\prime} \mathrm{B} \tag {2}
\]

In the above equations, \(\mathrm{E}^{\prime}\mathrm{Y}\), \(\mathrm{E}^{\prime}\mathrm{R}\), \(\mathrm{E}^{\prime}\mathrm{G}\) and \(\mathrm{E}^{\prime}\mathrm{B}\) can take floating point value spawning the range [0.0, 1.0]. \(\mathrm{E}^{\prime}\mathrm{Y}\) represents the luma information (gamma-corrected luminance). Two color difference components are derived from \(\mathrm{E}^{\prime}\mathrm{Y}\):

\[
\mathrm{E} ^ {\prime} \mathrm{B} - \mathrm{E} ^ {\prime} \mathrm{Y} \tag {3}
\]

\[
\mathrm{E} _ {\mathrm{R}} ^ {\prime} - \mathrm{E} _ {\mathrm{Y}} ^ {\prime}
\]

To facilitate the notation, this convention defines R', G' and B' as the gamma-corrected full scale values of the RGB color components, while their counterpart r', g' and b' are the gamma-corrected scaled down values as per BT.601 and BT.709. This section provides equations for 8-bit per color component, but a similar reasoning can be established for 10-bit components (or other bit depths).

\[
\mathrm{R} ^ {\prime}, \mathrm{G} ^ {\prime} \text {   and   } \mathrm{B} ^ {\prime} \text {   in   the   range   [0,255]   (8 - bit) } \tag {4}
\]

\[
\mathrm{r} ^ {\prime}, \mathrm{g} ^ {\prime} \text {   and   } \mathrm{b} ^ {\prime} \text {   in   the   range   [16,235]   (8 - bit) }
\]

#### 8.2.1 Generic Full Scale Y'CbCr (8-bit)

The full scale Y'CbCr is derived from using the basic luma equation from BT.601 and by having Y', Cb and Cr occupy the full 8-bit range of possible values. This format is not covered by BT.601 or BT.709, but often used in computer systems.