|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

#### 2.3.4 GenICam Control

Contains the features related to GenICam control and access of a specific Device module.

Table 2-13: GenICam Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  Root | M | All | ICategory | R | - | B | Provides the Root of the GenICam features tree.  |
|  DevicePort | M | All | IPort | R/W | - | I | The GenICam port through which the Device module is accessed.  |

#### 2.3.5 Event Control

Category that contains Event Control features.

Table 2-14: Event Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  EventControl | R | All | ICategory | R | - | E | Category that contains Event control features.  |
|  EventSelector | R | All | IEnumeration | R/W | - | E | Selects which Event to signal to the host application.  |
|  EventNotification[EventSelector] | R | All | IEnumeration | R/W | - | E | Activate or deactivate the notification to the host application of the occurrence of the selected Event.  |