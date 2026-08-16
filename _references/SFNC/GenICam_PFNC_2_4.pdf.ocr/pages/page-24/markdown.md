|  GEN<i>CAM |   | ![img-28.jpeg](img-28.jpeg) emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

#### 3.1.12 Polarized Location

Filters array variants are explicitly listed in this specification to reduce the number of possible pixel formats and thus overall complexity. If new sensors with a different filter array structure are released this filter array definition must be added to this specification. Every new filter array will result in a new subsection in this document.

##### Naming the formats

Naming is prefixed by POL (for Polarized) followed by the filter kernel size (x and y) followed by the type (M for monochrome or C for color) and a number if there are several different variations possible. The format's name will be listed in the FPNC-Specification within the following chapters.

In order to allow for a X and/or Y mirroring of the components there can be an added X for flip at X axis and/or a Y for flip at Y axis.

POL22M1: Polarized filter array 2x2 pixels according to POL22M1 below

POL22M1X: As above but filter array is reversed by x.

POL22M1Y: As above but filter array is reversed by y.

POL22M1XY: As above but filter array is reversed by x and y.

##### POL22M1 Location

POL22M1 defines a 2 by 2 filter array with no color filter applied.

|  90 | 45  |
| --- | --- |
|  135 | 0  |

Figure 3-24 POL22M1 array

Example of a complete code with 8 bits per pixel:

POL22M1_8: As defined in above figure.

POL22M1XY_8: Reversed by x and y