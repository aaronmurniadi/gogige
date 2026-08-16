|  **GEN<i>CAM** |   |   |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  an3dExtractionSelector] |  |  |  |  |  | transform of a point from the current (Anchor or Transformed) system to the reference system.  |
| --- | --- | --- | --- | --- | --- | --- |
|  Scan3dCoordinateReferenceValue[Scan3dExtractionSelector][Scan3dCoordinateReferenceSelector] | O | IFloat | R | - | E | Returns the reference value selected.  |
|  Scan3dFocalLength[RegionSelector] | O | IFloat | R | Pixel | E | Returns the focal length of the camera in pixel.  |
|  Scan3dBaseline | O | IFloat | R | m | E | Returns the baseline as the physical distance of two cameras in a stereo camera setup.  |
|  Scan3dPrincipalPointU[RegionSelector] | O | IFloat | R | Pixel | E | Returns the value of the horizontal position of the principal point, relative to the region origin, i.  |
|  Scan3dPrincipalPointV[RegionSelector] | O | IFloat | R | Pixel | E | Returns the value of the vertical position of the principal point, relative to the region origin, i.  |

## 2.20Light Control

Contains the features related to the Lighting Control (See the Light Control chapter for details).

Table 2-20: Lighting Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  LightControl | O | ICategory | R | - | B | Category containing the Lighting control features.  |
|  LightControllerSelector | O | IEnumeration | R/W | - | B | Selects the Light Controller to configure.  |
|  LightControllerSource[LightControllerSelector] | O | IEnumeration | R/W | - | B | Selects the input source signal of the Light Controller.  |
|  LightCurrentRating[LightControllerSelector] | O | IFloat | R/W | Amp | B | Set the current rating of the lighting output.  |
|  LightVoltageRating[LightControllerSelector] | O | IFloat | R/W | Volt | B | Set the voltage rating of the lighting output.  |
|  LightBrightness[LightControllerSelector] | O | IFloat | R/W | % | B | Set the brightness of the lighting output in percent.  |