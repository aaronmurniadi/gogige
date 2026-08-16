|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

- **B: Beginner** – Features that should be visible for *all* users via the GUI and API. This is the default visibility in the GenICam XML files and will be used if the Visibility element is omitted for a feature. The number of features with "Beginner" visibility should be limited to all **basic** features of the GenTL Producer so the GUI display is well-organized and is easy to use.
- **E: Expert** – Features that require a more in-depth knowledge of the device functionality. This is the preferred visibility level for all advanced features in the devices.
- **G: Guru** – Advanced features that might bring the devices into a state where it will not work properly anymore if it is set incorrectly for the devices current mode of operation.
- **I: Invisible** – Features that should be kept hidden for the GUI users but still be available via the API.

This document lists for each feature, a recommended Visibility that should be used.

### Selector

A selector is used to index which instance of the feature is accessed in situations where multiple instances of a feature exist.

A selector is a separate feature that is typically an IEnumeration or an IInteger. Selectors must be used only to select the target features for subsequent changes. It is not allowed to change the behavior of a GenTL Producer in response to a change of a selector value.

If a selector has only one possible value, the selector relation can be omitted but it is recommended to leave the selector feature as read only for information purpose.

In this document, the features which potentially dependent on a selector are expressed using the C language convention for arrays: a pair of brackets follows the feature name, like in SelectedFeature[Selector]. When the Selector is not present, one must deduce the feature is not an array.