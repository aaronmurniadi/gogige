Table 2-8: GenICam Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  Root | M | All | ICategory | R | - | B | Provides the Root of the GenICam features tree.  |
|  InterfacePort | M | All | IPort | R/W | - | I | The GenICam port through which the Interface module is accessed.  |

#### 2.2.5 Event Control

Category that contains Event Control features.

Table 2-9: Event Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  EventControl | R | All | ICategory | R | - | E | Category that contains Event control features.  |
|  EventSelector | R | All | IEnumeration | R/W | - | E | Selects which Event to signal to the host application.  |
|  EventNotification[EventSelector] | R | All | IEnumeration | R/W | - | E | Activate or deactivate the notification to the host application of the occurrence of the selected Event.  |

### 2.3 Device Module

#### 2.3.1 Device Information

Contains the features related to general information about a specific Device module.

Table 2-10: Device Information Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |