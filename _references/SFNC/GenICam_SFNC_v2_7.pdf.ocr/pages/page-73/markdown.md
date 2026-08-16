![img-9.jpeg](img-9.jpeg)

|  tor] |  |  |  |  |  |   |
| --- | --- | --- | --- | --- | --- | --- |
|  FocalLength[OpticControllerSelector] | O | IFloat | R/W | mm | B | Focal length in millimeters (mm).  |
|  FocalLengthStepper[OpticControllerSelector] | O | IInteger | R/W | - | E | FocalLengthStepper controls the stepper value of the focal length.  |
|  ShutterInitialize[OpticControllerSelector] | O | ICommand | (R)/W | - | B | Initializes the shutter and makes it ready for use.  |
|  ShutterStatus[OpticControllerSelector] | O | IEnumeration | R | - | B | Reads the status of the shutter.  |
|  Shutter[OpticControllerSelector] | O | IInteger | R/W | - | B | Controls whether the shutter is opened or closed.  |
|  FilterInitialize[OpticControllerSelector] | O | ICommand | (R)/W | - | B | Initializes the filter and makes it ready for use.  |
|  FilterStatus[OpticControllerSelector] | O | IEnumeration | R | - | B | Reads the status of the filter.  |
|  Filter[OpticControllerSelector] | O | IInteger | R/W | - | B | Filter positions in native number system.  |
|  StabilizationInitialize[OpticControllerSelector] | O | ICommand | (R)/W | - | B | Initializes the stabilization and makes it ready for use.  |
|  StabilizationStatus[OpticControllerSelector] | O | IEnumeration | R | - | B | Reads the status of the stabilization.  |
|  Stabilization[OpticControllerSelector] | O | IInteger | R/W | - | B | Control of image stabilization function build into the optic controller.  |
|  MagnificationInitialize[OpticControllerSelector] | O | ICommand | (R)/W | - | B | Initializes the magnification and makes it ready for use.  |
|  MagnificationStatus[OpticControllerSelector] | O | IEnumeration | R | - | B | Reads the status of the magnification.  |
|  Magnification[OpticControllerSelector] | O | IFloat | R/W | - | B | Magnification of the lens, defined as ratio between apparent size of an image and its true size.  |
|  MagnificationStepper[OpticControllerSelector] | O | IInteger | R/W | - | E | MagnificationStepper controls the stepper value of the magnification.  |

## 2.22 Chunk Data Control

Contains the features related to the Chunk Data Control (See the Chunk Data Control chapter for details).