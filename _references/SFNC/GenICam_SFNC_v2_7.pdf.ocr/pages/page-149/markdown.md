|  GEN<i>CAM |   | ![img-44.jpeg](img-44.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

• Bpp20: 20 bits per pixel.
• Bpp24: 24 bits per pixel.
• Bpp30: 30 bits per pixel.
• Bpp32: 32 bits per pixel.
• Bpp36: 36 bits per pixel.
• Bpp48: 48 bits per pixel.
• Bpp64: 64 bits per pixel.
• Bpp96: 96 bits per pixel.

4.40 PixelColorFilter

|  Name | PixelColorFilter[ComponentSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | NoneBayerRGBayerGBBayerGRBayerBG  |

Type of color filter that is applied to the image.

This value must always be coherent with the PixelFormat feature.

Possible values are:

• None: No color filter.
• BayerRG: Bayer Red Green filter.
• BayerGB: Bayer Green Blue filter.
• BayerGR: Bayer Green Red filter.
• BayerBG: Bayer Blue Green filter.

4.41 PixelDynamicRangeMin

|  Name | PixelDynamicRangeMin[ComponentSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |