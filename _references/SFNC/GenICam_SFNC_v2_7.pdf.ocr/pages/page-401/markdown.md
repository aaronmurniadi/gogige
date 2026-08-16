|  GEN<i>CAM |   | ![img-133.jpeg](img-133.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- DisparityC_Linescan: Disparity 2.5D Depth map with varying B sampling. The distance is inversely proportional to the pixel (disparity) value. The B (Y) axis comes from the encoder chunk value.

|  Scan3dOutputMode | 3dScanType |   | PixelFormats | Comment  |
| --- | --- | --- | --- | --- |
|   |  Line | Area  |   |   |
|  UncalibratedC | X | X | Coord3D_C |   |
|  CalibratedABC_Grid | X | X | Coord3D_ABC | Can be sent as 3 planar parts A/B/C.  |
|  CalibratedABC_PointCloud | X | X | Coord3D_ABC | Can be sent as 3 planar parts A/B/C.  |
|  CalibratedAC | X |  | Coord3D_AC | The B (scan direction) scaling from scale and offset data.  |
|  CalibratedAC_Linescan | X |  | Coord3D_AC | The B (scan direction) scaling is in Encoder Chunk data.  |
|  CalibratedC | X | X | Coord3D_C | Only calibrated range.  |
|  CalibratedC_Linescan | X |  | Coord3D_C | Only calibrated range. The B (scan direction) scaling is in Encoder Chunk data.  |
|  RectifiedC | X | X | Coord3D_C | Typically used together with Cartesian coordinate system.  |
|  RectifiedC_Linescan | X |  | Coord3D_C | The B (scan direction) scaling in Encoder Chunk data. Typically used together with Cartesian coordinate system.  |
|  DisparityC | X | X | Coord3D_C |   |
|  DisparityC_Linescan | X |  | Coord3D_C | The B (scan direction) scaling in Encoder Chunk data.  |

#### Linescan modes:

There are a number of linescan specific modes, i.e. with the suffix _Linescan. These use embedded chunk encoder information to specify the displacement between the lines in the 3D data stream. These modes are only applicable in situations where the encoder scaling is known and the encoder values transmitted as chunk data for each scan line.