|  GEN<i>CAM |   | ![img-14.jpeg](img-14.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

3.2.1.1 InterfaceInformation

|  Name | InterfaceInformation  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains all Interface Information features of the Interface module.

3.2.1.2 InterfaceID

|  Name | InterfaceID  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

GenTL Producer wide unique identifier of the selected interface.

Corresponds to the INTERFACE_INFO_ID command of IFGetInfo function.

3.2.1.3 InterfaceDisplayName

|  Name | InterfaceDisplayName  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Recommended  |