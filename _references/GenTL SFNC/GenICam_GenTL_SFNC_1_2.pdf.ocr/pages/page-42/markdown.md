|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Visibility | Expert  |
| --- | --- |
|  Values | Any NULL-terminated string  |

Unique identifier of the GenTL Producer like a GUID.

Corresponds to the TL_INFO_ID command of TLGetInfo function.

##### 3.1.1.3 TLVendorName

|  Name | TLVendorName  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Name of the GenTL Producer vendor.

Corresponds to the TL_INFO_VENDOR command of TLGetInfo function.

##### 3.1.1.4 TLModelName

|  Name | TLModelName  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Name of the GenTL Producer to distinguish different kinds of GenTL Producer implementations from one vendor.