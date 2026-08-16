|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 27.2.1 TransportLayerControl

|  Name | TransportLayerControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains the transport Layer control features.

### 27.2.2 TLParamsLocked

|  Name | TLParamsLocked  |
| --- | --- |
|  Category | TransportLayerControl  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | -  |

Used by the Transport Layer to prevent critical features from changing during acquisition.

Possible values are:

- 0: No features are locked.
- 1: Transport Layer and Device critical features are locked and cannot be changed.

### 27.2.3 TLParamsLockedSelector

|  Name | TLParamsLockedSelector  |
| --- | --- |