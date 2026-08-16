## 2.1.4 Event Control

Category that contains Event Control features.

Table 2-4: Event Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  EventControl | R | All | ICategory | R | - | E | Category that contains Event control features.  |
|  EventSelector | R | All | IEnumeration | R/W | - | E | Selects which Event to signal to the host application.  |
|  EventNotification[EventSelector] | R | All | IEnumeration | R/W | - | E | Activate or deactivate the notification to the host application of the occurrence of the selected Event.  |