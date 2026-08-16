|  Scan3dExtractionMethod[Scan3dExtractionSelector] | O | IEnumeration | R/W | - | E | Selects the method for extracting 3D from the input sensor data.  |
| --- | --- | --- | --- | --- | --- | --- |
|  Scan3dDistanceUnit[Scan3dExtractionSelector] | O | IEnumeration | R/(W) | - | E | Specifies the unit used when delivering (calibrated) distance data.  |
|  Scan3dCoordinateSystem[Scan3dExtractionSelector] | O | IEnumeration | R/(W) | - | E | Specifies the Coordinate system to use for the device.  |
|  Scan3dOutputMode[Scan3dExtractionSelector] | O | IEnumeration | R/(W) | - | E | Controls the Calibration and data organization of the device and the coordinates transmitted.  |
|  Scan3dCoordinateSystemReference[Scan3dExtractionSelector] | O | IEnumeration | R/W | - | E | Defines coordinate system reference location.  |
|  Scan3dCoordinateSelector[Scan3dExtractionSelector] | O | IEnumeration | R/W | - | E | Selects the individual coordinates in the vectors for 3D information/transformation.  |
|  Scan3dCoordinateScale[Scan3dExtractionSelector][Scan3dCoordinateSelector] | O | IFloat | R/(W) | - | E | Scale factor when transforming a pixel from relative coordinates to world coordinates.  |
|  Scan3dCoordinateOffset[Scan3dExtractionSelector][Scan3dCoordinateSelector] | O | IFloat | R/(W) | - | E | Offset when transforming a pixel from relative coordinates to world coordinates.  |
|  Scan3dInvalidDataFlag[Scan3dExtractionSelector][Scan3dCoordinateSelector] | O | IBoolean | R/(W) | - | E | Enables the definition of a non-valid flag value in the data stream.  |
|  Scan3dInvalidDataValue[Scan3dExtractionSelector][Scan3dCoordinateSelector] | O | IFloat | R/(W) | - | E | Value which identifies a non-valid pixel if Scan3dInvalidDataFlag is enabled.  |
|  Scan3dAxisMin[Scan3dExtractionSelector][Scan3dCoordinateSelector] | O | IFloat | R | - | E | Minimum valid transmitted coordinate value of the selected Axis.  |
|  Scan3dAxisMax[Scan3dExtractionSelector][Scan3dCoordinateSelector] | O | IFloat | R | - | E | Maximum valid transmitted coordinate value of the selected Axis.  |
|  Scan3dCoordinateTransformSelector[Scan3dExtractionSelector] | O | IEnumeration | R/W | - | E | Sets the index to read/write a coordinate transform value.  |
|  Scan3dTransformValue[Scan3dExtractionSelector][Scan3dCoordinateTransformSelector] | O | IFloat | R/W | - | E | Specifies the transform value selected.  |
|  Scan3dCoordinateReferenceSelector[Sc | O | IEnumeration | R/W | - | E | Sets the index to read a coordinate system reference value defining the  |