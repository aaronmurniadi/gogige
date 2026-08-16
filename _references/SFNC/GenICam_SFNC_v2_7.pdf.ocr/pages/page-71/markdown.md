Version 2.7.1

Standard Features Naming Convention

|  LightConnectionStatus[LightControllerSelector] | O | IEnumeration | R | - | B | Status of a light connected to the controller's output Line.  |
| --- | --- | --- | --- | --- | --- | --- |
|  LightCurrentMeasured[LightControllerSelector] | O | IFloat | R | Amp | B | The measured current applied to the lighting.  |
|  LightVoltageMeasured[LightControllerSelector] | O | IFloat | R | Volt | B | The measured voltage applied to the lighting.  |

### 2.21 Optic Control

Contains the features related to the Optic Control (See the Optic Control chapter for details).

Table 2-21: Optic Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  OpticControl | O | ICategory | R | - | B | Category for optical control features.  |
|  OpticControllerSelector | O | IEnumeration | R/W | - | B | Selects which optic controller to configure.  |
|  OpticControllerInitialize[OpticControllerSelector] | O | ICommand | (R)/W | - | B | Initializes the optic controller and makes it ready for use.  |
|  OpticControllerDisconnect[OpticControllerSelector] | O | ICommand | (R)/W | - | B | Closes the optic controller connection and prepares it for removal.  |
|  OpticControllerAbort[OpticControllerSelector] | O | ICommand | (R)/W | - | B | Aborts the current command or feature access.  |
|  OpticControllerStatus[OpticControllerSelector] | O | IEnumeration | R | - | B | Reads the status of the optic controller.  |
|  OpticControllerVendorName[OpticControllerSelector] | O | IString | R | - | B | Name of the manufacturer of the optic controller.  |
|  OpticControllerFamilyName[OpticControllerSelector] | O | IString | R | - | B | Name of the device family of the optic controller.  |
|  OpticControllerModelName[OpticControllerSelector] | O | IString | R | - | B | Model name of the optic controller.  |