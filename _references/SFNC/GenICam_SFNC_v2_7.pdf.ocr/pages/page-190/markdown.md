### 5.6.2 TriggerMode

|  Name | TriggerMode[TriggerSelector]  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Off On  |

Controls if the selected trigger is active.

Possible values are:

- Off: Disables the selected trigger.
- On: Enable the selected trigger.

### 5.6.3 TriggerSoftware

|  Name | TriggerSoftware[TriggerSelector]  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Generates an internal trigger. TriggerSource must be set to Software.

### 5.6.4 TriggerSource

|  Name | TriggerSource[TriggerSelector]  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |