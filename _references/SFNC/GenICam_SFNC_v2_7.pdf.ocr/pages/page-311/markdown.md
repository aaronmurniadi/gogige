|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- UserSet0: Selects the user set 0.
- UserSet1: Selects the user set 1.
• ...

When Default User Set is selected and loaded using UserSetLoad, the device must be in default factory settings state and must make sure the continuous acquisition use case works directly. Default User Set is read-only and cannot be modified.

### 16.3 UserSetDescription

|  Name | UserSetDescription[UserSetSelector]  |
| --- | --- |
|  Category | UserSetControl  |
|  Level | Optional  |
|  Interface | IString  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Description of the selected User Set content.

The recommended factory default value for an unused UserSet is an empty string.

### 16.4 UserSetLoad

|  Name | UserSetLoad[UserSetSelector]  |
| --- | --- |
|  Category | UserSetControl  |
|  Level | Recommended  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Loads the User Set specified by UserSetSelector to the device and makes it active.

### 16.5 UserSetSave

|  Name | UserSetSave[UserSetSelector]  |
| --- | --- |