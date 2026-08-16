|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.3.1.7 DeviceFamilyName

|  Name | DeviceFamilyName  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Name of the product family of the remote device model.

Corresponds to the "DeviceFamilyName" feature of the remote device and is usually retrieved via the bootstrap register of the remote device.

### 3.3.1.8 DeviceVersion

|  Name | DeviceVersion  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

The version of the remote device model.

Corresponds to the "DeviceVersion" feature of the remote device and is usually retrieved via the bootstrap register of the remote device.

Corresponds to the DEVICE_INFO_VERSION command of DevGetInfo function.