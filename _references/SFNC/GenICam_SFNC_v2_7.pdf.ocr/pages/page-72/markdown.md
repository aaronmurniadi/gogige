|  OpticControllerSerialNumber[OpticControllerSelector] | O | IString | R | - | B | Serial number of the optic controller.  |
| --- | --- | --- | --- | --- | --- | --- |
|  OpticControllerVersion[OpticControllerSelector] | O | IString | R | - | B | Version of the optic controller.  |
|  OpticControllerFirmwareVersion[OpticControllerSelector] | O | IString | R | - | B | Version of the firmware in the optic controller.  |
|  OpticControllerTemperature[OpticControllerSelector] | O | IFloat | R | C | B | Temperature of the optic controller in degrees Celsius (C).  |
|  ApertureInitialize[OpticControllerSelector] | O | ICommand | (R)/W | - | B | Initializes the aperture and makes it ready for use.  |
|  ApertureStatus[OpticControllerSelector] | O | IEnumeration | R | - | B | Reads the status of the aperture.  |
|  Aperture[OpticControllerSelector] | O | IFloat | R/W | - | B | Sets the aperture (also called iris, f-number, f-stop or f/#) of the lens.  |
|  ApertureStepper[OpticControllerSelector] | O | IIInteger | R/W | - | E | ApertureStepper controls the stepper value of the aperture.  |
|  NumericalAperture[OpticControllerSelector] | O | IFloat | R/W | - | B | Sets the numerical aperture of a lens.  |
|  FocusInitialize[OpticControllerSelector] | O | ICommand | (R)/W | - | B | Initializes the focus and makes it ready for use.  |
|  FocusStatus[OpticControllerSelector] | O | IEnumeration | R | - | B | Reads the status of the focus.  |
|  FocusStepper[OpticControllerSelector] | O | IIInteger | R/W | - | E | FocusStepper controls the stepper value of the focus.  |
|  FocusAutoMode[OpticControllerSelector] | O | IEnumeration | R/W | - | B | Sets automatic focus mode.  |
|  FocusAuto [OpticControllerSelector] | O | IEnumeration | R/W | - | B | Sets automatic focus.  |
|  FocalPower [OpticControllerSelector] | O | IFloat | R/W | dpt | B | Sets the focal power (in diopters/dpt).  |
|  ObjectSensorDistance[OpticControllerSelector] | O | IFloat | R/W | mm | B | Distance from the image sensor surface to the object in millimeters (mm).  |
|  FocalLengthInitialize[OpticControllerSelector] | O | ICommand | (R)/W | - | B | Initializes the focal length and makes it ready for use.  |
|  FocalLengthStatus[OpticControllerSelector] | O | IEnumeration | R | - | B | Reads the status of the focal length.  |