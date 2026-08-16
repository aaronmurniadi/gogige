|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

For 8-bit data, the range of values for each component is:

❖ Y'601, r', g' and b' in range [16, 235], unsigned (220 levels)
❖ Cb and Cr in range [16, 240], signed shifted by 128, with 128 representing 0 (225 levels)

Values must be truncated to fit in that range.

The reverse equations are given by:

$$\mathrm{r}' = \quad \mathrm{Y}'_{601} \quad + 1.37071 \ (\mathrm{Cr} - 128)$$

$$\mathrm{g}' = \quad \mathrm{Y}'_{601} \quad - 0.33645 \ (\mathrm{Cb} - 128) \quad - 0.69820 \ (\mathrm{Cr} - 128)$$

$$\mathrm{b}' = \quad \mathrm{Y}'_{601} \quad + 1.73245 \ (\mathrm{Cb} - 128)$$

with Y'601 in the range [16, 235] and, Cb and Cr in the range [16, 240].

Equation 7: Y'CbCr601 to r'g'b' conversion (8 bits)

### 8.2.3 Y'CbCr709 (8-bit)

ITU-R BT.709 provides a definition of the Y', Cb and Cr based on (1). It defines signal range identical to BT.601, as expressed in (9).

But its luma equation is based on (2). Hence its 2 color difference signals are given by:

$$\mathrm{E}'_{\mathrm{B}} - \mathrm{E}'_{\mathrm{Y}} \quad \text{is in the range } [-0.9278, +0.9278] \tag{12}$$

$$\mathrm{E}'_{\mathrm{R}} - \mathrm{E}'_{\mathrm{Y}} \quad \text{is in the range } [-0.7874, +0.7874]$$

After normalization to occupy the [-0.5, 0.5] range:

$$\mathrm{E}'_{\mathrm{Cb}} = (0.5 / 0.9278) (\mathrm{E}'_{\mathrm{B}} - \mathrm{E}'_{\mathrm{Y}}) \tag{13}$$

$$\mathrm{E}'_{\mathrm{Cr}} = (0.5 / 0.7874) (\mathrm{E}'_{\mathrm{R}} - \mathrm{E}'_{\mathrm{Y}})$$

Considering (9) that provides the range for Y', Cb and Cr, (13) leads to:

$$\mathrm{Y}'_{709} = 219 \ \mathrm{E}'_{\mathrm{Y}} + 16 \tag{14}$$

$$\mathrm{Cb} = 224 \ \mathrm{E}'_{\mathrm{Cb}} + 128 = 224 \times (0.5 / 0.9278) (\mathrm{E}'_{\mathrm{B}} - \mathrm{E}'_{\mathrm{Y}}) + 128$$

$$\mathrm{Cr} = 224 \ \mathrm{E}'_{\mathrm{Cr}} + 128 = 224 \times (0.5 / 0.7874) (\mathrm{E}'_{\mathrm{R}} - \mathrm{E}'_{\mathrm{Y}}) + 128$$

Again, two options exist depending on the allowed range of values for the RGB components.