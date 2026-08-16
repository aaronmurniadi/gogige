|  ![img-110.jpeg](img-110.jpeg) GEN<i>CAM |   | ![img-111.jpeg](img-111.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Intensity | 5 | 2 | Region 0 | 2D image | Mono8  |
| --- | --- | --- | --- | --- | --- |
|  Chunk data | (6) | - | All components chunk data if enabled (not shown). | (Chunk data) | GenICam Chunk  |

If the setup is changed to have a data output with a 3D point cloud planar image by replacing:

ComponentSelector[Source1] = Disparity; // Disparity
ComponentEnable[Source1][Disparity] = True;
PixelFormat[Source1][Disparity] = Coord3D_C16; // 3D disparity output format.

by

ComponentSelector[Source1] = Range; // Range
ComponentEnable[Source1][Range] = True;
PixelFormat[Source1][Range] = Coord3D_ABC16_Planar; // 3D point cloud output format.

formatting could be defined as:

|  Component | Part | Source | Region | data type | data format  |
| --- | --- | --- | --- | --- | --- |
|  Range | 0 | 1 | Region 0 | 3D plane of tri planar image | Coord3D_A16  |
|  Range | 1 | 1 | Region 0 | 3D plane of tri planar image | Coord3D_B16  |
|  Range | 2 | 1 | Region 0 | 3D plane of tri planar image | Coord3D_C16  |
|  Confidence | 3 | 1 | Region 0 | 3D confidence image | Confidence8  |
|  Intensity | 4 | 1 | Region 0 | 2D image | Mono8  |
|  Range | 5 | 2 | Region 0 | 3D plane of tri planar image | Coord3D_A16  |
|  Range | 6 | 2 | Region 0 | 3D plane of tri planar image | Coord3D_B16  |
|  Range | 7 | 2 | Region 0 | 3D plane of tri planar image | Coord3D_C16  |
|  Confidence | 8 | 2 | Region 0 | 3D confidence image | Confidence8  |
|  Intensity | 9 | 2 | Region 0 | 2D image | Mono8  |
|  Chunk data | 10 | - | All components chunk data if enabled (not shown). | (Chunk data) | GenICam Chunk  |