|  ![img-235.jpeg](img-235.jpeg)CAN |   | ![img-236.jpeg](img-236.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Value | Description  |
| --- | --- | --- |
|   |  | complete image.The complete image consists of 3 planes.The planes of a given planar image must be placed as consecutive parts within the buffer.It is recommended to usePIXELFORMAT NAMESPACE PFNC 32BITdata format with this data type whenever possible.  |
|  PART_DATATYPE_3D_PLANE_QUADPLANAR | 8 | Single plane of a planar 3D image.The data should be linked with the other coordinate planes to get the complete image.The complete image consists of 4 planes.The planes of a given planar image must be placed as consecutive parts within the buffer.It is recommended to usePIXELFORMAT NAMESPACE PFNC 32BITdata format with this data type whenever possible.  |
|  PART_DATATYPE_CONFIDENCE_MAP | 9 | Confidence of the individual pixel values. Expresses the level of validity of given pixel values.Confidence map is always used together with one or more additional image-based parts matching 1:1 dimension-wise. Each value in the confidence map expresses level of validity of the image pixel at matching position.The data format should be a Confidence PFNC format.It is recommended to usePIXELFORMAT NAMESPACE PFNC 32BITdata format with this data type whenever possible.  |