|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

User-programmable device identifier of the remote device.

Corresponds to the “DeviceUserID” feature of the remote device and usually it is retrieved via the bootstrap register of the remote device.

Corresponds to the DEVICE_INFO_USER_DEFINED_NAME command of DevGetInfo function.

### 3.3.1.5 DeviceVendorName

|  Name | DeviceVendorName  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Name of the remote device vendor.

Corresponds to the DEVICE_INFO_VENDOR command of DevGetInfo function.

### 3.3.1.6 DeviceModelName

|  Name | DeviceModelName  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Name of the remote device model.

Corresponds to the DEVICE_INFO_MODEL command of DevGetInfo function.