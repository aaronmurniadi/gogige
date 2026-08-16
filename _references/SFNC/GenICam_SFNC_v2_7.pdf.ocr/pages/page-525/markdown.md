|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 27.4.12 GevInterfaceSelector

|  Name | GevInterfaceSelector  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥0  |

Selects which logical link to control.

Note: The number of physical network interfaces may be greater than the value reported by GevInterfaceSelector. This is generally the case when link aggregation is activated.

### 27.4.13 GevLinkSpeed (Deprecated)

|  Name | GevLinkSpeed[GevInterfaceSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Mbps  |
|  Visibility | Invisible  |
|  Values | >0  |

This feature is deprecated (See DeviceLinkSpeed). It was representing the speed of transmission negotiated by the given logical link.

### 27.4.14 GevMACAddress

|  Name | GevMACAddress[GevInterfaceSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |