|  Component designation | Positioning in Image | Description  |
| --- | --- | --- |
|  “BayerRG” | Bayer_LMMN location | Bayer filter Red-Green-Green-Blue  |
|  “BayerGB” | Bayer_MNLM location | Bayer filter Green-Blue-Red-Green  |
|  “BayerBG” | Bayer_NMML location | Bayer filter Blue-Green-Green-Red  |
|  “BiColorRGBG” | BiColor_LMNO location | Bi-color pixel Red/Green - Blue/Green  |
|  “BiColorGRGB” | BiColor_LMNO location | Bi-color pixel Green/Red - Green/Blue  |
|  “BiColorBGRG” | BiColor_LMNO location | Bi-color pixel Blue/Green - Red/Green  |
|  “BiColorGBGR” | BiColor_LMNO location | Bi-color pixel Green/Blue - Green/Red  |
|  “aRGB” | LMNO4444 location | alpha-Red-Green-Blue alpha component content is manufacturer-specific.  |
|  “YRGB” | LMNO4444 location | Luma-Red-Green-Blue  |
|  “RGBa” | LMNO4444 location | Red-Green-Blue-alpha alpha component content is manufacturer-specific.  |
|  “aBGR” | LMNO4444 location | alpha-Blue-Green-Red alpha component content is manufacturer-specific.  |
|  “BGRa” | LMNO4444 location | Blue-Green-Red-alpha alpha component content is manufacturer-specific.  |
|  “YUV” “YUV422” “YUV411” | LMN444 location LMN422 location LMN411 location | YUV color space, typically an incorrect usage of the Y’CbCr color space. Legacy from IIDC standard. *Default:* Y is unsigned, U and V are signed (shifted by adding 128 for 8-bit components)  |
|  “YCbCr” “YCbCr422” “YCbCr411” | LMN444 location LMN422 location LMN411 location | Generic Y’CbCr color space using full range of 256 values for each component. See section 8.2.1 for the color transform equations. Y’, Cb and Cr are in the range [0, 255]. Y is unsigned, Cb and Cr are signed (shifted by adding 128).  |
|  “YCbCr601” “YCbCr601_422” “YCbCr601_411” | LMN444 location LMN422 location LMN411 location | Y’CbCr color space as specified by ITU-R BT.601 (SDTV). See section 8.2.2 for the color transform equations. Y’ is in the range [16, 235]. Cb and Cr are in the range [16, 240]. Y’ is unsigned, Cb and Cr are signed (shifted by adding 128).  |
|  “YCbCr709” “YCbCr709_422” “YCbCr709_411” | LMN444 location LMN422 location LMN411 location | Y’CbCr color space as specified by ITU-R BT.709 (HDTV). See section 8.2.3 for the color transform equations. Y’ is in the range [16, 235]. Cb and Cr are in the range [16, 240]. Y’ is unsigned, Cb and Cr are signed (shifted by adding 128).  |
|  “SCF1WBWG” | SCF1_LMLN location | Sparse Color Filter #1, White-Blue-White-Green pattern  |