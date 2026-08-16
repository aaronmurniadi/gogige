|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

The multi-slope exposure control section describes all features related to controlling the multiple phases of the exposure of the photosensitive cells during image acquisition.

To define a generic interface to an HDR feature also known as multi-slope or multi-knee-point or piecewise linear response, two slightly different models are considered:

- The pixels are reset to a specific level at a certain point in time.
- The pixel capacity is controlled and can be changed at a certain point in time.

Additionally, in both implementations, partial exposure times may be controlled separately.

Both models result in one image with a non-linear exposure characteristic composed of multiple, piecewise linear slopes. The points in graphs where the linear segments meet are generally named "knee-points".

To convert the parameters of the sensors to features, the following was considered:

- The number of knee-points differs from sensor to sensor.
- The number of used knee-points may be configured.
- All sensors allow adjusting the partial exposure times, but not necessarily the levels/capacities.
- Multi-slope exposure for HDR-like images requires different times for each part of the whole exposure.
- As a generalization, the threshold for a knee-point is determined in "percent" (of the currently available pixel capacity or of the currently available reset voltage - The currently available values might be influenced by the camera gain settings).
- As a further generalization, the unit for the sub-exposures is also "percent" (of the current exposure time).

For an explanation of the parameters involved in multi-slope exposure, two knee-points are assumed: