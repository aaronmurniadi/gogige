|  ![img-27.jpeg](img-27.jpeg) |   | ![img-28.jpeg](img-28.jpeg)  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- Scatter: The acquisition of data measuring how much light is scattered around the reflected light. In processing this is used as an additional intensity image, often together with the standard intensity or reflectance.
- Disparity: The acquisition of stereo camera disparity data is controlled. Disparity is a more specific range format approximately inversely proportional to distance. Disparity is typically given in pixel units.
- Multispectral: The acquisition of multiple spectral bands corresponding to various light wavelengths is controlled.

# Multiple components from a region:

The ComponentSelector allows to control the multiple components of a single source region. In this case, all components are defined to be registered (i.e. for each pixel in all components, the data in all images refer to the same source region position.

Although the source data region for all the components is the same, the resulting components might be of different size. In that case, ComponentSelector can be used to specify that and for example, Width[SourceSelector][RegionSelector][ComponentSelector] would give the Size X of each individual component. Another example could be that each component would have a different PixelFormat (e.g. floating point for Range data, unsigned integer for Intensity and single bit for Confidence).

Example of multiple components acquisition from a sensor.

// Setup for 3D Range+Confidence+Reflectance sent as calibrated Range map.
Scan3dOutputMode = CalibratedABC_Grid; // XYZ Planar output.
ComponentSelector = Range;
ComponentEnable[Range] = True;
ComponentSelector = Confidence;
ComponentEnable[Confidence] = True;
ComponentSelector = Reflectance;
ComponentEnable[Reflectance] = True;