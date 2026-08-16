|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.1.2.1 InterfaceEnumeration

|  Name | InterfaceEnumeration  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains all Interface Enumeration features of the System module.

### 3.1.2.2 InterfaceUpdateList

|  Name | InterfaceUpdateList  |
| --- | --- |
|  Category | InterfaceEnumeration  |
|  Level | Mandatory  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Updates the internal list of the interfaces. This feature should be readable if the execution cannot performed immediately. The command then returns and the status can be polled. This function interacts with the TLUpdateInterfaceList function of the GenTL Producer. It is up to the GenTL Consumer to handle access in case both methods are used.

### 3.1.2.3 InterfaceUpdateTimeout

|  Name | InterfaceUpdateTimeout  |
| --- | --- |
|  Category | InterfaceEnumeration  |
|  Level | Recommended  |
|  Interface | IInteger  |