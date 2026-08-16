|  ColorTransformationControl | R | ICategory | R | - | E | Category that contains the Color Transformation control features.  |
| --- | --- | --- | --- | --- | --- | --- |
|  ColorTransformationSelector | O | IEnumeration | R/W | - | E | Selects which Color Transformation module is controlled by the various Color Transformation features.  |
|  ColorTransformationEnable[ColorTransformationSelector] | O | IBoolean | R/W | - | E | Activates the selected Color Transformation module.  |
|  ColorTransformationValueSelector[ColorTransformationSelector] | O | IEnumeration | R/W | - | E | Selects the Gain factor or Offset of the Transformation matrix to access in the selected Color Transformation module.  |
|  ColorTransformationValue[ColorTransformationSelector][ColorTransformationValueSelector] | O | IFloat | R/W | - | E | Represents the value of the selected Gain factor or Offset inside the Transformation matrix.  |

## 2.7 Digital I/O Control

Contains the features related to the control of the input and output pins of the device (See the Digital I/O Control chapter for details).

Table 2-7: Digital I/O Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  DigitalIOControl | R | ICategory | R | - | E | Category that contains the digital input and output control features.  |
|  LineSelector | R | IEnumeration | R/W | - | E | Selects the physical line (or pin) of the external device connector or the virtual line of the Transport Layer to configure.  |
|  LineMode[LineSelector] | O | IEnumeration | R/W | - | E | Controls if the physical Line is used to Input or Output a signal.  |
|  LineInverter[LineSelector] | R | IBoolean | R/W | - | E | Controls the inversion of the signal of the selected input or output Line.  |
|  LineStatus[LineSelector] | R | IBoolean | R | - | E | Returns the current status of the selected input or output Line.  |
|  LineStatusAll | O | IInteger | R | - | E | Returns the current status of all available Line signals at time of polling in a single bitfield.  |
|  LineSource[LineSelector] | R | IEnumeration | R/W | - | E | Selects which internal acquisition or I/O source signal to output on the selected Line.  |
|  LineFormat[LineSelector] | O | IEnumeration | R/W | - | E | Controls the current electrical format of the selected physical input or output Line.  |