|  ![img-233.jpeg](img-233.jpeg)CAN |   | ![img-234.jpeg](img-234.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  PART_DATATYPE_2D_PLANE_QUADPLANAR | 4 | Single color plane of a planar (2D) image. The data should be linked with the other color planes to get the complete image.The complete image consists of 4 planes.The planes of a given planar image must be placed as consecutive parts within the buffer.It is recommended to usePIXELFORMAT NAMESPACE PFNC 32BITdata format with this data type whenever possible.  |
|  PART_DATATYPE_3D_IMAGE | 5 | 3D image (pixel coordinates). This part carries all the pixel data of given image (even if the image is represented by a single-plane pixel format, for example when transferring the depth map only).It is recommended to usePIXELFORMAT NAMESPACE PFNC 32BITdata format with this data type whenever possible.  |
|  PART_DATATYPE_3D_PLANE_BIPLANAR | 6 | Single plane of a planar 3D image. The data should be linked with the other coordinate planes to get the complete image.The complete image consists of 2 planes.The planes of a given planar image must be placed as consecutive parts within the buffer.It is recommended to usePIXELFORMAT NAMESPACE PFNC 32BITdata format with this data type whenever possible.  |
|  PART_DATATYPE_3D_PLANE_TRIPLANAR | 7 | Single plane of a planar 3D image. The data should be linked with the other coordinate planes to get the  |