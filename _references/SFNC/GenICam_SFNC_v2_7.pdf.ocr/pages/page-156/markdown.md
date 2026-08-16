|  GEN<i>CAM |   | ![img-48.jpeg](img-48.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Visibility | Expert  |
| --- | --- |
|  Values | Device-specific  |

Control the rate of the produced compressed stream.

This feature is available when ImageCompressionRateOption is equals to FixBitrate or if the device only supports the FixBitrate mode.

The list of valid values is device specific.

4.47.5 ImageCompressionJPEGFormatOption

|  Name | ImageCompressionJPEGFormatOption  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | LosslessBaselineStandardBaselineOptimizedProgressiveDevice-specific  |

When JPEG is selected as the compression format, a device might optionally offer better control over JPEG-specific options through this feature.

Possible values are:

- Lossless: Selects lossless JPEG compression based on a predictive coding model.
- BaselineStandard: Indicates this is a baseline sequential (single-scan) DCT-based JPEG.
- BaselineOptimized: Provides optimized color and slightly better compression than baseline standard by using custom Huffman tables optimized after statistical analysis of the image content.
- Progressive: Indicates this is a progressive (multi-scan) DCT-based JPEG.