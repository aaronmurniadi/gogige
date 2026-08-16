|  Access | Read/Write  |
| --- | --- |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Intensity Infrared Ultraviolet Range Reflectance Confidence Scatter Disparity Multispectral Device-specific  |

Selects a component to activate/deactivate its data streaming.

The **ComponentSelector** determines which data acquisition component to control. In a standard (2D) Areascan or Linescan device, Intensity is the typical image component available and it can typically not be turned off. However, for 3D cameras it is often relevant to be able to enable and disable the acquisition of different components.

In a general case, different regions of the sensor can be used to acquire and transmit different content. A typical case is found in many linescan 3D cameras where a region is used to acquire 3D Components while another is used to acquire a different 2D Component. In that case, the ComponentEnable feature will be selected by a RegionSelector in addition to the usual ComponentSelector.

Note that this use case can also be addressed using a SourceSelector and defining multiple (virtual) sensors as described in the Source Control chapter.

Possible values are:

- **Intensity**: The acquisition of intensity (monochrome or color) of the visible reflected light is controlled.
- **Infrared**: The acquisition of non-visible infrared light is controlled.
- **Ultraviolet**: The acquisition of non-visible ultraviolet light is controlled.
- **Range**: The acquisition of range (distance) data is controlled. The data produced may be only range (2.5D) or a point cloud giving the 3D coordinates depending on the Scan3dControl features.
- **Reflectance**: The reflected intensity acquired together with Range in a Linescan3D sensor acquiring a single linescan profile for each exposure of the sensor.
- **Confidence**: The acquisition of confidence map of the acquired image is controlled. Confidence data may be binary (0 - invalid) or an integer where 0 is invalid and increasing value is increased confidence in the data in the corresponding pixel. If floating point representation is used the confidence image is normalized to the range [0,1], for integer representation the maximum possible integer represents maximum confidence.