|  ![img-34.jpeg](img-34.jpeg) |   | ![img-35.jpeg](img-35.jpeg)  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- Disparity: The acquisition of stereo camera disparity data is controlled. Disparity is a more specific range format approximately inversely proportional to distance. Disparity is typically given in pixel units.

4.17 ImageComponentEnable (Deprecated)

|  Name | ImageComponentEnable[RegionSelector][ComponentSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | TrueFalse  |

This feature is deprecated (See ComponentEnable). It was used to control if the selected component streaming is active.

To help backward compatibility, this feature can be included as Invisible in the device's XML.

4.18Width

|  Name | Width[RegionSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | >0  |

Width of the image provided by the device (in pixels).

This reflects the current Region of interest. The maximum value of this feature takes into account horizontal binning, decimation, or any other function changing the maximum horizontal dimensions of the image and is typically equal to WidthMax - OffsetX.

This feature is generally mandatory for transmitters and transceivers of most Transport Layers.

4.19 Height

|  Name | Height[RegionSelector]  |
| --- | --- |