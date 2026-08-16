|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  EventTestData | R | ICategory | R | - | E | Category that contains all the data features related to the Event Test generated using the TestEventGenerate command.  |
| --- | --- | --- | --- | --- | --- | --- |
|  EventTest | R | IInteger | R | - | E | Returns the unique identifier of the Event Test type of event generated using the TestEventGenerate command.  |
|  EventTestTimestamp | R | IInteger | R | - | E | Returns the Timestamp of the Event Test event.  |

## 2.14 User Set Control

Contains the features related to the User Set Control to save and load the user device settings (See the User Set Control chapter for details).

Table 2-14: User Set Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  UserSetControl | R | ICategory | R | - | B | Category that contains the User Set control features.  |
|  UserSetSelector | R | IEnumeration | R/W | - | B | Selects the feature User Set to load, save or configure.  |
|  UserSetDescription[UserSetSelector] | O | IString | R/W | - | B | Description of the selected User Set content.  |
|  UserSetLoad[UserSetSelector] | R | ICommand | (R)/W | - | B | Loads the User Set specified by UserSetSelector to the device and makes it active.  |
|  UserSetSave[UserSetSelector] | R | ICommand | (R)/W | - | B | Save the User Set specified by UserSetSelector to the non-volatile memory of the device.  |
|  UserSetDefault | O | IEnumeration | R/W | - | B | Selects the feature User Set to load and make active by default when the device is reset.  |
|  UserSetDefaultSelector | O | IEnumeration | R/W | - | I | This feature is deprecated (See UserSetDefault).  |
|  UserSetFeatureSelector | R | IEnumeration | R/W | - | E | Selects which individual UserSet feature to control.  |
|  UserSetFeatureEnable[UserSetFeatureSelector] | R | IBoolean | R/(W) | - | E | Enables the selected feature and make it active in all the UserSets.  |

## 2.15 Sequencer Control

Contains the features related to the control of the Sequencer for features change (See the Sequencer Control chapter for details).