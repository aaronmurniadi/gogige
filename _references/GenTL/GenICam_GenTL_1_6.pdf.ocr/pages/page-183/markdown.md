|  GEN<I>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Name | Interface | Access | Description  |
| --- | --- | --- | --- |
|   |  |  | DWRITE. • “ReadOnly” The device is available to be opened with read-only access. As soon as the device is open the value should change to “OpenRead”. Corresponds to DEVICE_ACCESS_STATUS_READONLY. • “NoAccess” The device is seen by the producer but not reachable. Corresponds to DEVICE_ACCESS_STATUS_NOACCESS. • “Busy” The device is already opened by another entity. Corresponds to DEVICE_ACCESS_STATUS_BUSY. • “OpenReadWrite” The device is already open by this GenTL Producer with RW access. Corresponds to DEVICE_ACCESS_STATUS_OPEN_READWRITE. • “OpenReadOnly” The device is already opened by this GenTL Producer with RO access. Corresponds to DEVICE_ACCESS_STATUS_OPEN_READONLY.  |

### 7.1.3 Device Module

The Device module contains all features which must be accessible in the Device module: Port functions use the DEV_HANDLE to access these features. The Port access for this module is mandatory.

Do not mistake this Device module Port access with the remote device Port access. This module represents the GenTL Producer’s view on the remote device. The remote device port is retrieved via the DevGetPort function returning a PORT_HANDLE for the remote device.

Table 7-9: Device information features