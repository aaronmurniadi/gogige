|  **GEN<i>CAM** |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

## 5 Optional “Data type” indicator

This field allows to specify the data type of the pixel component values. If omitted, unsigned integer data type is assumed.

|  *Empty* or “u” | Default for most component. Unsigned integer data.  |
| --- | --- |
|  “s” | Signed integer data (two’s complement).  |
|  “f” | Floating point data (binary floating point format compatible with IEC 60559:1989 standard) *Note: need to handle specific floating point values, in particular NaN’s, during pixel data processing might incur performance penalties, it might be desirable to avoid such values within pixel data whenever possible.*  |

Use one value if all components have same data type, otherwise must list as many data type indicators as there are components in the pixel, successively with no space in-between in the same order they are presented in the “Components and Location” field.