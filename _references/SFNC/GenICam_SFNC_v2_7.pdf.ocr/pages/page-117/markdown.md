Each pixel in the image has a format defined by the PixelFormat feature. Only a subset of the possible pixel formats is presented in this document. The complete list of possible standard pixel formats and their layout can be found in the separate "GenICam Pixel Format Naming Convention (PFNC)" specification (See the GenICam download page on the EMVA web site). This web page also gives the list of the currently standardized Pixel Formats and their unique Identifier value (See the "GenICam Pixel Format Values" and "Reference header file for PFNC" documents).

Because the PixelFormat feature contains a mix of informations specified by the user and informations provided by the device, it is suitable for describing the whole pixel settings but might be less practical when individual setting must be set or inquired. Therefore a second set of features exists composed of the individual components of PixelFormat. Those features are PixelSize, PixelColorFilter, PixelDynamicRangeMin and PixelDynamicRangeMax.

Even if the PixelFormat might allow for, e.g. 16 bits per pixel, the real image data might provide only a certain range of value (e.g. 12 bits per pixel because the camera is equipped with a 12 bit analog to digital converter only). In that case, PixelDynamicRangeMin and PixelDynamicRangeMax specify the lower and upper limits of the pixel values in the image. In general, PixelDynamicRangeMin should be zero and PixelDynamicRangeMax should be a power of two ($$[0, 2^{DataDepth} - 1]$$). There should be no missing codes in the range.

### 4.1 ImageFormatControl

|  Name | ImageFormatControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category for Image Format Control features.

### 4.2 SensorWidth

|  Name | SensorWidth  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |