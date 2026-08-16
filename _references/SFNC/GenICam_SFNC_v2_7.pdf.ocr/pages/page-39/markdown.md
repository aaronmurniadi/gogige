|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- B: Beginner - Features that should be visible for all users via the GUI and API. This is the default visibility in the GenICam XML files and will be used if the Visibility element is omitted for a feature. The number of features with "Beginner" visibility should be limited to all basic features of the devices so the GUI display is well-organized and is easy to use.
- E: Expert - Features that require a more in-depth knowledge of the camera functionality. This is the preferred visibility level for all advanced features in the cameras.
- G: Guru – Advanced features that might bring the cameras into a state where it will not work properly anymore if it is set incorrectly for the cameras current mode of operation.
- I: Invisible – Features that should be kept hidden for the GUI users but still be available via the API.

This document lists for each feature, a Visibility that should be used.

## Selector

A selector is used to index which instance of the feature is accessed in situations where multiple instances of a feature exist (For instance, the analog gain for each separate channel of the red/green/blue component of a color camera).

A selector is a separate feature that is typically an IEnumeration or an IInteger. Selectors must be used only to select the target features for subsequent changes. It is not allowed to change the behavior of a device in response to a change of a selector value.

If a selector has only one possible value, the selector relation can be omitted but it is recommended to leave the selector feature as read only for information purpose (Ex: TriggerSelector = FrameStart (read only) for a device that has only this trigger type supported).

In this document, the features potentially dependent on a selector are expressed using the C language convention for arrays: a pair of brackets follows the feature name, like in SelectedFeature[Selector]. When the Selector is not present, one must deduce the feature is not an array.

In general, a selector should apply only to a single category of feature (Ex: TriggerSelector applies only to the Trigger related features). However, it is possible that certain more advanced devices will require a selector that applies to features in different categories. For example a device with 2 independent input sensors could have a SourceSelector feature that would select features in the Image Format Control, Acquisition Control, Analog Control, LUT Control and Color Transformation Control categories in order to globally control all the features associated with a particular source (Ex": SourceSelector = Source1, PixelFormat[SourceSelector] = Mono8, Gain[SourceSelector] = 10, AcquisitionStart[SourceSelector]).

Note also that when a feature that has a selector is persisted to a file, the selector is iterated to allow saving the complete array of values and not only the currently selected element.