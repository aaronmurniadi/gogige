|  Visibility | Expert  |
| --- | --- |
|  Values | FixBitrate FixQuality ... Device-specific  |

Two rate controlling options are offered: fixed bit rate or fixed quality. The exact implementation to achieve one or the other is vendor-specific.

Note that not all compression techniques or implementations may support this feature.

Possible values are:

- FixBitrate: Output stream follows a constant bit rate. Allows easy bandwidth management on the link.
- FixQuality: Output stream has a constant image quality. Can be used when image processing algorithms are sensitive to image degradation caused by excessive data compression.

### 4.47.3 ImageCompressionQuality

|  Name | ImageCompressionQuality  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | Integer  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Device-specific  |

Control the quality of the produced compressed stream.

This feature is available when ImageCompressionRateOption is equal to FixQuality or if the device only supports the FixQuality mode.

The list of valid values is device-specific. A higher value means a better quality for the produced compressed stream.

### 4.47.4 ImageCompressionBitrate

|  Name | ImageCompressionBitrate  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | Float  |
|  Access | Read/(Write)  |
|  Unit | Mbps  |