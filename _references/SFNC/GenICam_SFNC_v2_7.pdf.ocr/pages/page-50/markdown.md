![img-4.jpeg](img-4.jpeg)

|  TimestampLatchValue | O | Integer | R | ns | E | Returns the latched value of the timestamp counter.  |
| --- | --- | --- | --- | --- | --- | --- |

## 2.2 Image Format Control

Contains the features related to the format of the transmitted image (See the Image Format Control chapter for details).

Table 2-2: Image Format Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  ImageFormatControl | R | ICategory | R | - | B | Category for Image Format Control features.  |
|  SensorWidth | R | IInteger | R | - | E | Effective width of the sensor in pixels.  |
|  SensorHeight | R | IInteger | R | - | E | Effective height of the sensor in pixels.  |
|  SensorPixelWidth | O | IFloat | R | um | G | Physical size (pitch) in the x direction of a photo sensitive pixel unit.  |
|  SensorPixelHeight | O | IFloat | R | um | G | Physical size (pitch) in the y direction of a photo sensitive pixel unit.  |
|  SensorName | O | IString | R | - | G | Product name of the imaging Sensor.  |
|  SensorShutterMode | O | IEnumeration | R/(W) | - | G | Specifies the shutter mode of the device.  |
|  SensorTaps | O | IEnumeration | R/(W) | - | E | Number of taps of the camera sensor.  |
|  SensorDigitizationTaps | O | IEnumeration | R/(W) | - | E | Number of digitized samples outputted simultaneously by the camera A/D conversion stage.  |
|  WidthMax | R | IInteger | R | - | E | Maximum width of the image (in pixels).  |
|  HeightMax | R | IInteger | R | - | E | Maximum height of the image (in pixels).  |
|  RegionSelector | O | IEnumeration | R/(W) | - | B | Selects the Region of interest to control.  |
|  RegionMode[RegionSelector] | O | IEnumeration | R/W | - | B | Controls if the selected Region of interest is active and streaming.  |
|  RegionDestination[RegionSelector] | O | IEnumeration | R/(W) | - | E | Control the destination of the selected region.  |
|  RegionIDValue[RegionSelector] | O | IInteger | R | - | E | Returns a unique Identifier value that corresponds to the selected Region.  |
|  ComponentSelector | O | IEnumeration | R/W | - | B | Selects a component to activate/deactivate its data streaming.  |