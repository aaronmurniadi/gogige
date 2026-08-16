|  GEN<i>CAM |   | ![img-47.jpeg](img-47.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- Off: The device doesn't perform de-interlacing.
- LineDuplication: The device performs de-interlacing by outputting each line of each field twice.
- Weave: The device performs de-interlacing by interleaving the lines of all fields.

### 4.47 Image Compression

This section describes the feature related to image compression.

#### 4.47.1 ImageCompressionMode

|  Name | ImageCompressionMode  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Visibility | Beginner  |
|  Unit | -  |
|  Values | OffJPEGJPEG2000H264...Device-specific  |

Enable a specific image compression mode as the base mode for image transfer.

Possible values are:

- Off: Default value. Image compression is disabled. Images are transmitted uncompressed.
• JPEG: JPEG compression is selected.
• JPEG2000: JPEG 2000 compression is selected.
• H264: H.264 compression is selected.

#### 4.47.2 ImageCompressionRateOption

|  Name | ImageCompressionRateOption  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |