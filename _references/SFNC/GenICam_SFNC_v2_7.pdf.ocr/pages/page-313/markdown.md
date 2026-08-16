|  Category | UserSetControl  |
| --- | --- |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | Default UserSet0 (if 0 based), UserSet1, ...  |

This feature is deprecated (See UserSetDefault). Selects the feature User Set to load and make active when the device is reset.

To help backward compatibility, this feature can be included as Invisible in the device's XML.

Possible values are:

- Default: Select the factory setting user set.
- UserSet0: Select the user set 0.
- UserSet1: Select the user set 1.
- ...

If Default is selected, the device will boot with the default factory settings and makes sure the continuous acquisition use case works directly.

### 16.8 UserSetFeatureSelector

|  Name | UserSetFeatureSelector  |
| --- | --- |
|  Category | UserSetControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Device-Specific-Feature-List  |

Selects which individual UserSet feature to control.

The feature lists all the features that can be a part of a device UserSet. All the device's UserSets have the same features.