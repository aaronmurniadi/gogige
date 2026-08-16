|  GEN<i>CAM |   | ![img-10.jpeg](img-10.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Level | Recommended  |
| --- | --- |
|  Interface | IString  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

User readable name of the GenTL Producer.

Corresponds to the TL_INFO_DISPLAYNAME command of TLGetInfo function.

3.1.1.8 TLPath

|  Name | TLPath  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Full path to the GenTL Producer including filename and extension.

Corresponds to the TL_INFO_PATHNAME command of TLGetInfo function.

3.1.1.9 TLType

|  Name | TLType  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Mandatory  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |