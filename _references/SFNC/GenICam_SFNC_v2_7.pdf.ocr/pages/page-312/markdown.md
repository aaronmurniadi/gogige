|  Category | UserSetControl  |
| --- | --- |
|  Level | Recommended  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Save the User Set specified by UserSetSelector to the non-volatile memory of the device.

### 16.6 UserSetDefault

|  Name | UserSetDefault  |
| --- | --- |
|  Category | UserSetControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Default UserSet0 (if 0 based), UserSet1, ...  |

Selects the feature User Set to load and make active by default when the device is reset.

Possible values are:

- Default: Select the factory setting user set.
- UserSet0: Select the user set 0.
- UserSet1: Select the user set 1.
- ...

If Default is selected, the device will boot with the default factory settings and makes sure the continuous acquisition use case is ready to be used.

### 16.7 UserSetDefaultSelector (Deprecated)

|  Name | UserSetDefaultSelector  |
| --- | --- |