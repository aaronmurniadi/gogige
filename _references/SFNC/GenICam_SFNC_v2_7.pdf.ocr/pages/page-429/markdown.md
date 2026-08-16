|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- NotInitialized: The optic controller is not initialized.
- NotSupported: The optic controller is physically connected but not supported.
- Busy: The optic controller executes a feature access/command.
- Ready: The optic controller is ready to use.
- Error: The optic controller encountered an error.

### 23.4.7 OpticControllerVendorName

|  Name | OpticControllerVendorName[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated String  |

Name of the manufacturer of the optic controller.

### 23.4.8 OpticControllerFamilyName

|  Name | OpticControllerFamilyName[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated String  |

Name of the device family of the optic controller.

### 23.4.9 OpticControllerModelName

|  Name | OpticControllerModelName[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |