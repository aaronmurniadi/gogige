|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

#### 3.2.1 CFA Basic Components

Pixel formats based on a color filter array must explicitly state the basic components used to create the pattern.

Table 3-2 : CFA Basic Components

|  Basic Component for CFA | Color | Additional Information  |
| --- | --- | --- |
|  “R” | Red | Used in primary color sensor.  |
|  “G” | Green | Used in primary and complementary color sensor.  |
|  “B” | Blue | Used in primary color sensor.  |
|  “W” | White | A pixel with no color filter (panchromatic)  |
|  “C” | Cyan | Used in complementary color sensor.  |
|  “M” | Magenta | Used in complementary color sensor.  |
|  “Ye” | Yellow | Used in complementary color sensor.  |
|  “Ir” | Infrared | Used for infrared (IR) channel  |

### 3.3 Generic Data Types Formats

This section provides the format naming to be used for generic non-pixel data. Those generic formats should be used only for non-image data block representation. They can be used for example, to specify the format of a metadata block of information, processing results (such as histogram) ...

Note that although strictly speaking those formats are not representing pixels, they are defined in this PFNC naming convention in order to keep all data formats used by GenICam in one document.

Table 3-3 : Generic Data Types Designation

|  Component designation | Positioning | Description  |
| --- | --- | --- |
|  “Data” | Same as Mono location | Generic data.  |