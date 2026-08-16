### 4.16 ImageComponentSelector (Deprecated)

|  Name | ImageComponentSelector  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | Intensity Color Infrared Ultraviolet Range Confidence Scatter Disparity Device-specific  |

This feature is deprecated (See **ComponentSelector**). It was used to select a component.

To help backward compatibility, this feature can be included as Invisible in the device's XML.

Possible values are:

- **Intensity**: The acquisition of intensity (monochrome or color) of the visible reflected light is controlled.
- **Color**: The acquisition of color of the reflected light is controlled
- **Infrared**: The acquisition of non-visible infrared light is controlled.
- **Ultraviolet**: The acquisition of non-visible ultraviolet light is controlled.
- **Range**: The acquisition of range (distance) data is controlled. The data produced may be only range (2.5D) or a point cloud 3D coordinates depending on the Scan3dControl features.
- **Confidence**: The acquisition of confidence map of the acquired image is controlled. Confidence data may be binary (0 - invalid) or an integer where 0 is invalid and increasing value is increased confidence in the data in the corresponding pixel. If floating point representation is used the confidence image is normalized to the range [0,1], for integer representation the maximum possible integer represents maximum confidence.
- **Scatter**: The acquisition of data measuring how much light is scattered around the reflected light. In processing this is used as an additional intensity image, often together with the standard intensity.