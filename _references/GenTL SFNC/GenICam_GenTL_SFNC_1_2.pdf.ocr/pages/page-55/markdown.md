|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Visibility | Expert  |
| --- | --- |
|  Values | InterfaceListChanged  |

Selects which Event to signal to the host application.

Possible values are:

- InterfaceListChanged: the list of interfaces is updated.

### 3.1.4.3 EventNotification

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
- Once: The selected Event notification is enabled for one event then return to the Off state.

## 3.2 Interface Module

Contains all features of the Interface module that are independent from the underlying transport technology.

### 3.2.1 Interface Information

Features in this section provide basic information about the Interface Module and its identity. Note that all features in this section are defined read-only.