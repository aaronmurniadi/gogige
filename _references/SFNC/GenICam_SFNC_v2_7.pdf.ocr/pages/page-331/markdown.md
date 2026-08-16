### 18.1 FileAccessControl

|  Name | FileAccessControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | -  |

Category that contains the File Access control features.

### 18.2 FileSelector

|  Name | FileSelector  |
| --- | --- |
|  Category | FileAccessControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | UserSetDefault UserSet1 UserSet2 UserSet3 ... LUTLuminance LUTRed LUTGreen LUTBlue ...  |

Selects the target file in the device.

The entries of this enumeration define the names of all files in the device that can be accessed via the File access.

Possible values are: