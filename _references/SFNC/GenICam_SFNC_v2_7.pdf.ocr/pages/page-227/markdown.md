|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|   | Offset2 is the blue offset  |
| --- | --- |
|   | C0_{out} is the first resulting component of the pixel after the transformation  |
|   | C1_{out} is the second resulting component of the pixel after the transformation  |
|   | C2_{out} is the third resulting component of the pixel after the transformation  |

Example for YUV conversion:

The Color Transformation can also be used outside of the simple scope of color correction on RGB pixels. For instance, it can be used as a color convert to convert RGB to YUV.

Here is the example of this conversion for 8-bit pixels:

$$\left( \begin{array}{c} Y \\ U \\ V \end{array} \right) = \left( \begin{array}{c c c} 0.299 & 0.587 & 0.114 \\ -0.147 & -0.289 & 0.436 \\ 0.615 & -0.515 & -0.100 \end{array} \right) \left( \begin{array}{c} R_{in} \\ G_{in} \\ B_{in} \end{array} \right) + \left( \begin{array}{c} 0 \\ 128 \\ 128 \end{array} \right)$$

### 8.1 ColorTransformationControl

|  Name | ColorTransformationControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains the Color Transformation control features.

### 8.2 ColorTransformationSelector

|  Name | ColorTransformationSelector  |
| --- | --- |
|  Category | ColorTransformationControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |