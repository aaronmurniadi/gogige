|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Specifies timeout for the DeviceUpdateList Command.

#### 3.2.2.4 DeviceSelector

|  Name | DeviceSelector  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Selector for the different devices on this interface. This value only changes on execution of "DeviceUpdateList". The selector is 0-based in order to match the index of the C interface.

#### 3.2.2.5 DeviceID

|  Name | DeviceID[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Interface wide unique identifier of the selected device. This value only changes on execution of the DeviceUpdateList command.

Corresponds to the IFGetDeviceID function with the index corresponding to "DeviceSelector".

#### 3.2.2.6 DeviceVendorName

|  Name | DeviceVendorName[DeviceSelector]  |
| --- | --- |