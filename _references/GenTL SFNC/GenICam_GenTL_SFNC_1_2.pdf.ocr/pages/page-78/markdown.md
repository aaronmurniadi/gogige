|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | EventControl  |
| --- | --- |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | InterfaceLost DeviceListChanged  |

Selects which Event to signal to the host application.

Possible values are:

- InterfaceLost: Raised when the interface connection is lost.
- DeviceListChanged: The list of devices is updated.

### 3.2.5.3 EventNotification

|  Name | EventNotification[EventSelector]  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off On Once  |

Activate or deactivate the notification to the host application of the occurrence of the selected Event.

Possible values are:

- Off: The selected Event notification is disabled.
- On: The selected Event notification is enabled.