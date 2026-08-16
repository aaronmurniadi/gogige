|  ![img-183.jpeg](img-183.jpeg) CAM |   | ![img-184.jpeg](img-184.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

#### 6.4.3 Device Enumerations

##### 6.4.3.1 DEVICE_ACCESS_FLAGS

enum DEVICE_ACCESS_FLAGS

This enumeration defines different modes how a device is to be opened with the IFOpenDevice function. The values can not be combined.

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  DEVICE_ACCESS_UNKNOWN | 0 | Not used in a command. It can be used to initialize a variable to query that information.  |
|  DEVICE_ACCESS_NONE | 1 | Deprecated: do not use.(This value was never usable in any meaningful way.)  |
|  DEVICE_ACCESS_READONLY | 2 | Opens the device read only. All Port functions can only read from the device.  |
|  DEVICE_ACCESS_CONTROL | 3 | Opens the device in a way that other hosts/processes can have read only access to the device. Device access level is read/write for this process.  |
|  DEVICE_ACCESS_EXCLUSIVE | 4 | Open the device in a way that only this host/process can have access to the device. Device access level is read/write for this process.  |
|  DEVICE_ACCESS_CUSTOM_ID | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific.If a generic GenTL Consumer is using custom DEVICE_ACCESS_FLAGSs provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |

##### 6.4.3.2 DEVICE_ACCESS_STATUS

enum DEVICE_ACCESS_STATUS

This enumeration defines the status codes used in the info functions with the info command DEVICE_INFO_ACCESS_STATUS to retrieve the current accessibility of the device.

|  Enumerator | Value | Description  |
| --- | --- | --- |