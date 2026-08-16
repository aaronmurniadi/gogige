|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Access | Read/Write  |
| --- | --- |
|  Unit | ms  |
|  Visibility | Expert  |
|  Values | >0  |

Specifies timeout for the InterfaceUpdateList Command.

### 3.1.2.4 InterfaceSelector

|  Name | InterfaceSelector  |
| --- | --- |
|  Category | InterfaceEnumeration  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥0  |

Selector for the different GenTL Producer interfaces. This interface list only changes on execution of "InterfaceUpdateList". The selector is 0-based in order to match the index of the C interface.

### 3.1.2.5 InterfaceID

|  Name | InterfaceID[InterfaceSelector]  |
| --- | --- |
|  Category | InterfaceEnumeration  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

GenTL Producer wide unique identifier of the selected interface.