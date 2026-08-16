## 16 User Set Control

This chapter describes the features for global control of the device settings. It allows loading or saving factory or user-defined settings.

Loading the factory default User Set guarantees a state where a continuous acquisition can be started using only the mandatory features.

### 16.1 UserSetControl

|  Name | UserSetControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains the User Set control features.

### 16.2 UserSetSelector

|  Name | UserSetSelector  |
| --- | --- |
|  Category | UserSetControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Default UserSet0 (if 0 based), UserSet1, ...  |

Selects the feature User Set to load, save or configure.

Possible values are:

- Default: Selects the factory setting user set.