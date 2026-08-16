### 3.11 DeviceID (Deprecated)

|  Name | DeviceID  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Recommended  |
|  Interface | ISString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | Any NULL-terminated string  |

This feature is deprecated (See DeviceSerialNumber). It was representing the Device unique identifier (serial number).

To help backward compatibility, this feature can be included as Invisible in the device's XML.

### 3.12 DeviceUserID

|  Name | DeviceUserID  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | ISString  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

User-programmable device identifier.

When this feature is present, it must be writable and should be persistent.

The recommended factory default value is an empty string.

### 3.13 DeviceSFNCVersionMajor

|  Name | DeviceSFNCVersionMajor  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Recommended  |
|  Interface | Integer  |
|  Access | Read  |
|  Unit | -  |