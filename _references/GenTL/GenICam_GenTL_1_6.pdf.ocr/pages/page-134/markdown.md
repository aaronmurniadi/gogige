|  ![img-185.jpeg](img-185.jpeg) CAM |   | ![img-186.jpeg](img-186.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  DEVICE_ACCESS_STATUS_UNKNOWN | 0 | The current availability of the device is unknown.  |
| --- | --- | --- |
|  DEVICE_ACCESS_STATUS_READWRITE | 1 | The device is available to be opened for Read/Write access but it is currently not opened. This value will only be returned through IFGetDeviceInfo function because as soon as the device is open DEVICE_ACCESS_STATUS_OPEN_READWRITE will be returned.  |
|  DEVICE_ACCESS_STATUS_READONLY | 2 | The device is available to be opened for Read access but is currently not opened. In case the device allows both read and write access the value DEVICE_ACCESS_STATUS_READWRITE has a higher priority. This value will only be returned through IFGetDeviceInfo function because as soon as the device is open DEVICE_ACCESS_STATUS_OPEN_READONLY will be returned.  |
|  DEVICE_ACCESS_STATUS_NOACCESS | 3 | The device is seen be the producer but is not available for access because it is not reachable.  |
|  DEVICE_ACCESS_STATUS_BUSY | 4 | The device is already owned/opened by another entity.  |
|  DEVICE_ACCESS_STATUS_OPEN_READWRITE | 5 | The device is already owned/opened by this GenTL Producer with RW access. A further call to IFOpenDevice will return GC_ERR_RESOURCE_IN_USE.  |
|  DEVICE_ACCESS_STATUS_OPEN_READONLY | 6 | The device is already owned/opened by this GenTL Producer with RO access. A further call to IFOpenDevice will return GC_ERR_RESOURCE_IN_USE.  |
|  DEVICE_ACCESS_STATUS_CUSTOM_ID | 1000 | Starting value for custom IDs which are implementation specific. If a generic GenTL Consumer is using custom DEVICE_ACCESS_STATUS ids provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |