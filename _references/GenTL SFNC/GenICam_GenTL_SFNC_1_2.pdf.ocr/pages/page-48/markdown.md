|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Interface | IInteger  |
| --- | --- |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |

This feature is deprecated (See InterfaceTLVersionMajor).

Major version number of the GigE Vision specification the GenTL Producer implementation complies with.

If the value of the feature TLType is "Mixed" but supports GigE Vision interfaces this feature must be present.

3.1.1.15 GevVersionMinor (Deprecated)

|  Name | GevVersionMinor  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Optional  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

This feature is deprecated (See InterfaceTLVersionMinor).

Minor version number of the GigE Vision specification the GenTL Producer implementation complies with.

If the value of the feature TLType is "Mixed" but supports GigE Vision interfaces this feature must be present.

#### 3.1.2 Interface Enumeration

The Interface Enumeration section describes all features related to discovery and enumeration of interfaces belonging to the System module.