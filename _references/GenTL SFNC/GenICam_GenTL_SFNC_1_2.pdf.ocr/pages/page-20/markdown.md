## 2 Features Summary

This section provides a comprehensive list of the standard features covered by this document. The following sections provide more detailed explanation of each feature.

### 2.1 System Module

#### 2.1.1 System Information

Contains the features related to general information about the GenTL Producer.

Table 2-1: System Information Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  SystemInformation | R | All | ICategory | R | - | B | Category that contains all System Information features of the System module.  |
|  TLID | M | All | IString | R | - | E | Unique identifier of the GenTL Producer like a GUID.  |
|  TLVendorName | M | All | IString | R | - | B | Name of the GenTL Producer vendor.  |
|  TLModelName | M | All | IString | R | - | B | Name of the GenTL Producer to distinguish different kinds of GenTL Producer implementations from one vendor.  |
|  TLVersion | M | All | IString | R | - | B | Vendor specific version string of the GenTL Producer.  |
|  TLFileName | R | All | IString | R | - | E | Filename including extension of the GenTL Producer.  |
|  TLDisplayName | R | All | IString | R/(W) | - | B | User readable name of the GenTL Producer.  |
|  TLPath | M | All | IString | R | - | E | Full path to the GenTL Producer including filename and extension.  |
|  TLType | M | All | IEnumeration | R | - | E | Transport layer type of the GenTL Producer implementation.  |