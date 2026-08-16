|  **GEN<i>CAM** |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

Ex: YCbCr420_8_YY_CbCr_Semiplanar . Here we have a plane/component sequence YY_CbCr with 2 planes separated by ‘_’, one YY plane and one CbCr plane.

### 7.3 Components Sequencing

If the component sequence is not identical to the one specified in the “Components and Location” field, then the actual sequence should be provided by listing as many components as necessary to correctly determine the correct sequence in image memory.

For instance, there are various sequences of components for Y’CbCr, some sending Y’, followed by Cb and Cr. But there are others where Cr is put first. This is further complicated by sub-sampling of the chroma components. In these cases, it might be necessary to unequivocally list the order that the components are transmitted.

In this case, a underscore (‘_’) is put before listing the sequence of component to clearly separate the list from the rest of the pixel name.

Ex: YUV411_8_UYYVYY