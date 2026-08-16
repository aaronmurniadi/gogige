|  **GEN<i>CAM** |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

## 6.7 Dealing with Line and Image Boundaries

The packing styles in this section do not specify how it is affected at a line or image boundary. Two options are allowed when data is put in a PFNC-compliant image buffer:

1. **Image Padding:** no artificial padding is inserted at the end of a line. Hence the first pixel of a given line might not start on an 8-bit boundary as it might be combined in the same byte as the last pixel of the previous line. At the end of the image, missing luma components from a cluster take a value of 0
2. **Line Padding:** the last pixel of each line is padded to complete on an 8-bit boundary (or to the boundary specified by the standard referencing this convention), so the first pixel of the next line starts on a fresh 8-bit boundary. At the end of the line, missing luma components from a cluster take a value of 0

**Important:** The standard referencing this Pixel Format Naming Convention is expected to explicitly define in their document the method used when dealing with line and image boundaries for data put in a PFNC-compliant image buffer. A special treatment might be necessary for 3D data types.

As an example, assume a Mono1p image where the image width is not a multiple of 8. The last pixel of the first line does not align to an 8-bit boundary and a choice must be made between image padding, where pixels from different lines might be packed/grouped together, and line padding where pixels from different lines are not packed/grouped together.

**Note:** The last pixel at the image or line boundary might need special considerations when grouped or packed style is used. If a component is missing to complete the packing group, then one or more additional 'artificial' components with a value of 0 must be used.

For instance, assume the pixel format uses a cluster of 3 luma components with line padding, but only 2 are left at the end of the line. In this case, an extra luma with a value of 0 must be used to complete this cluster to enable the packing.

**Note:** Special care might be required when working with pixel data which are not line oriented (such as a 3D point cloud). In such case line padding might not apply. If there could be a risk for any confusion, the camera might elect to prefer pixel formats that ensure each pixel starts at an 8-bit boundary.