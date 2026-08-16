|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Access | Read  |
| --- | --- |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains all Device Enumeration features of the Interface module.

### 3.2.2.2 DeviceUpdateList

|  Name | DeviceUpdateList  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Updates the internal device list. This feature should be readable if the execution cannot be performed immediately. The command then returns and the status can be polled. This feature interacts with the IFUpdateDeviceList function of the GenTL Producer. It is up to the GenTL Consumer to handle access in case both methods are used.

### 3.2.2.3 DeviceUpdateTimeout

|  Name | DeviceUpdateTimeout  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | ms  |
|  Visibility | Expert  |
|  Values | >0  |