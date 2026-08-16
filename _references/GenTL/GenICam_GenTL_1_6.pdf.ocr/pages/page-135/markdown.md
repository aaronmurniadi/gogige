|  GENICAM |   | ![img-187.jpeg](img-187.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.4.3.3 DEVICE_INFO_CMD

enum DEVICE_INFO_CMD

This enumeration defines commands to retrieve information with the DevGetInfo function on a device handle or with IFGetDeviceInfo. The reported information using these two functions should be in sync if the information is available. This is also true for the info command DEVICE_INFO_ACCESS_STATUS.

The column labeled “Impl” in the following table lists if the implementation of a given command is mandatory (M), optional (O) or conditional mandatory (CM). Mandatory means that a GenTL Producer must implement the listed command. Optional means that it is up to the implementor if a given command is implemented or not. Conditional Mandatory means that command is to be implemented if possible.

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|  DEVICE_INFO_ID | M | 0 | Unique ID of the device.Data type: STRING  |
|  DEVICE_INFO_VENDOR | M | 1 | Device vendor name.Data type: STRING  |
|  DEVICE_INFO_MODEL | M | 2 | Device model name.Data type: STRING  |
|  DEVICE_INFO_TLTYPE | M | 3 | Transport layer technology that is supported. See string constants in chapter 6.6.1.Data type: STRING  |
|  DEVICE_INFO_DISPLAYNAME | M | 4 | User readable name of the device. If this is not defined in the device this should be “VENDOR MODEL (ID)”.Data type: STRING  |
|  DEVICE_INFO_ACCESS_STATUS | O | 5 | Gets the access status the GenTL Producer has on the device.Data type: INT32(DEVICE_ACCESS_STATUS enumeration value)  |
|  DEVICE_INFO_USER_DEFINED_NAME | O | 6 | String containing the user defined name of the device. If the information is not available, the query should result inGC_ERR_NOT_AVAILABLE.Data type: STRING  |
|  DEVICE_INFO_SERIAL_NUMBER | CM | 7 | Serial number of the device in string format. If the information is not available,the query should result inGC_ERR_NOT_AVAILABLE.Data type: STRING  |