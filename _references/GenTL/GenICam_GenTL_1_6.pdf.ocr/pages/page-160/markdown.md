|  ![img-231.jpeg](img-231.jpeg)CAN |   | ![img-232.jpeg](img-232.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Value | Description  |
| --- | --- | --- |
|   |  | the data type of the data in the provided buffer part. From the GenTL Consumer perspective this can be handled as raw data.  |
|  PART_DATATYPE_2D_IMAGE | 1 | Color or monochrome (2D) image. This part carries all the pixel data of given image (even if the image is represented by a single-plane pixel format). It is recommended to use PIXELFORMAT_NAMESPACE_PFN_C_32BIT data format with this data type whenever possible.  |
|  PART_DATATYPE_2D_PLANE_BIPLANAR | 2 | Single color plane of a planar (2D) image. The data should be linked with the other color planes to get the complete image. The complete image consists of 2 planes. The planes of a given planar image must be placed as consecutive parts within the buffer. It is recommended to use PIXELFORMAT_NAMESPACE_PFN_C_32BIT data format with this data type whenever possible.  |
|  PART_DATATYPE_2D_PLANE_TRIPLANAR | 3 | Single color plane of a planar (2D) image. The data should be linked with the other color planes to get the complete image. The complete image consists of 3 planes. The planes of a given planar image must be placed as consecutive parts within the buffer. It is recommended to use PIXELFORMAT_NAMESPACE_PFN_C_32BIT data format with this data type whenever possible.  |