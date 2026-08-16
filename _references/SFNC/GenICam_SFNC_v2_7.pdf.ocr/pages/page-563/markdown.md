|  Interface | IBoolean  |
| --- | --- |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Enable or disable the link sharing functionality of the device.

Note: All CxpLinkSharing features support Link Sharing as defined in the CXP specification.

### 27.7.6 CxpLinkSharingSubDeviceSelector

|  Name | CxpLinkSharingSubDeviceSelector  |
| --- | --- |
|  Category | CoaXPress  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Index of the sub device used in the Link Sharing.

Index 0 is the master sub device, other index are for slave sub devices.

### 27.7.7 CxpLinkSharingStatus

|  Name | CxpLinkSharingStatus[CxpLinkSharingSubDeviceSelector]  |
| --- | --- |
|  Category | CoaXPress  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |