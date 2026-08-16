### 3.3 DeviceScanType

|  Name | DeviceScanType  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Areascan Linescan Areascan3D Linescan3D  |

Scan type of the sensor of the device.

Typically, this feature is not writable. But some cameras might allow switching between linescan and areascan.

Possible values are:

- Areascan: 2D sensor outputting an image created from a unique sensor acquisition.
- Linescan: 1D sensor outputting an image acquired line by line.
- Areascan3D: 3D sensor outputting a range (or disparity) image created from a unique sensor acquisition.
- Linescan3D: 3D sensor outputting a range (or disparity) image acquired line by line.

### 3.4 DeviceVendorName

|  Name | DeviceVendorName  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Name of the manufacturer of the device.