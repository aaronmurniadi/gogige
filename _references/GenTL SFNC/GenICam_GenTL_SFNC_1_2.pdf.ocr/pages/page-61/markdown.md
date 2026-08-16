|  GEN<i>CAM |   | ![img-15.jpeg](img-15.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Interface | IInteger  |
| --- | --- |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

IP address of the selected subnet of this interface. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

3.2.1.12 GevInterfaceSubnetMask

|  Name | GevInterfaceSubnetMask[GevInterfaceSubnetSelector]  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Subnet mask of the selected subnet of this interface. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

#### 3.2.2 Device Enumeration

The Device Enumeration section describes all features related to discovery and enumeration of devices belonging to the Interface module.

3.2.2.1 DeviceEnumeration

|  Name | DeviceEnumeration  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |