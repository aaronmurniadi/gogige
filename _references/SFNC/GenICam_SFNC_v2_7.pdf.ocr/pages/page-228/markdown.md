|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Visibility | Expert  |
| --- | --- |
|  Values | RGBtoRGB RGBtoYUV Device-specific  |

Selects which Color Transformation module is controlled by the various Color Transformation features.

Possible values are:

- RGBtoRGB: RGB to RGB color transformation.
- RGBtoYUV: RGB to YUV color transformation.

It is typically not available when a single Color Transformation module is supported.

### 8.3 ColorTransformationEnable

|  Name | ColorTransformationEnable[ColorTransformationSelector]  |
| --- | --- |
|  Category | ColorTransformationControl  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Activates the selected Color Transformation module.

### 8.4 ColorTransformationValueSelector

|  Name | ColorTransformationValueSelector[ColorTransformationSelector]  |
| --- | --- |
|  Category | ColorTransformationControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |