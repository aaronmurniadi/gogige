|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Visibility | Expert  |
| --- | --- |
|  Values | UncalibratedC CalibratedABC_Grid CalibratedABC_PointCloud CalibratedAC CalibratedAC_Linescan CalibratedC CalibratedC_Linescan RectifiedC RectifiedC_Linescan DisparityC DisparityC_Linescan  |

Returns the Calibrated Mode of the payload image.

Possible values are:

- UncalibratedC: Uncalibrated 2.5D Depth map. The distance data does not represent a physical unit and may be non-linear. The data is a 2.5D range map only.
- CalibratedABC_Grid: 3 Coordinates in grid organization. The full 3 coordinate data with the grid array organization from the sensor kept.
- CalibratedABC_PointCloud: 3 Coordinates without organization. The full 3 coordinate data without any organization of data points. Typically only valid points transmitted giving varying image size.
- CalibratedAC: 2 Coordinates with fixed B sampling. The data is sent as a A and C coordinates (X,Z or Theta,Rho). The B (Y) axis uses the scale and offset parameters for the B axis.
- CalibratedAC_Linescan: 2 Coordinates with varying sampling. The data is sent as a A and C coordinates (X,Z or Theta,Rho). The B (Y) axis comes from the encoder chunk value.
- CalibratedC: Calibrated 2.5D Depth map. The distance data is expressed in the chosen distance unit. The data is a 2.5D range map. No information on X-Y axes available.
- CalibratedC_Linescan: Depth Map with varying B sampling. The distance data is expressed in the chosen distance unit. The data is a 2.5D range map. The B (Y) axis comes from the encoder chunk value.
- RectifiedC: Rectified 2.5D Depth map. The distance data has been rectified to a uniform sampling pattern in the X and Y direction. The data is a 2.5D range map only. If a complete 3D point cloud is rectified but transmitted as explicit coordinates it should be transmitted as one of the "CalibratedABC" formats.
- RectifiedC_Linescan: Rectified 2.5D Depth map with varying B sampling. The data is sent as rectified 1D profiles using Coord3D_C pixels. The B (Y) axis comes from the encoder chunk value.
- DisparityC: Disparity 2.5D Depth map. The distance is inversely proportional to the pixel (disparity) value.