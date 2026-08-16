|  ![img-30.jpeg](img-30.jpeg) |   | ![img-31.jpeg](img-31.jpeg)  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

4.13 ComponentIDValue

|  Name | ComponentIDValue[ComponentSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns a unique Identifier value that corresponds to type of the component selected by ComponentSelector.

This value is typically used by the Transport Layer to specify the type of information included in the transmitted component data.

The standard values that must be used for the known component types listed in the ComponentSelector feature are (range 0x0000-Ox7FFF):

|  Undefined | = 0 (Reserved)  |
| --- | --- |
|  Intensity | = 1  |
|  Infrared | = 2  |
|  Ultraviolet | = 3  |
|  Range | = 4  |
|  Reflectance | = 5  |
|  Confidence | = 6  |
|  Scatter | = 7  |
|  Disparity | = 8  |
|  Multispectral | = 9  |

Other standard Component types that can be used:

|  Metadata | = 0x8001  |
| --- | --- |
|  Custom | = 0xFF00-0xFFFE  |
|  Reserved | = 0xFFFF  |

Note: Custom components are device specific component types that can be used when none of the predefined types apply. Those custom Component types must have a unique ID in the range 0xFF00-0xFFFE.

4.14 GroupSelector

|  Name | GroupSelector  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |