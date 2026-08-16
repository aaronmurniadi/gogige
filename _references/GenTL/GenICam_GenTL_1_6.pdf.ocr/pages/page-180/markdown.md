|  ![img-270.jpeg](img-270.jpeg)CAN |   | ![img-271.jpeg](img-271.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

## 7 Standard Features Naming Convention for GenTL

The different GenTL modules expose their features through the Port functions interface. To interpret the virtual register map of each module the GenICam GenApi is used. This document only contains the names of mandatory features that must be implemented to guarantee interoperability between the different GenTL Consumers and GenTL Producers. Additional features and descriptions can be found in the GenICam Standard Features Naming Convention document (SFNC) and in the GenTL Standard Features Naming Convention document (GenTL SFNC).

For technical reasons the different transport layer technologies and protocols (e.g., GigE Vision, IIDC 1394, Camera Link, etc.) have different feature sets. This is addressed in dedicated sections specialized on these technologies. Also features specific to one technology have a prefix indicating its origin, e.g., Gev for GigE Vision specific features or CI for Camera Link specific features. Mixed-type GenTL Producers must implement mandatory features of all supported technologies in the System node map. The mandatory technology specific features falling under the “InterfaceSelector” might be marked not-available (NA) when an interface implementing other technology is currently selected.

Interface, Device, Data Stream and Buffer node maps are unequivocally bound to a particular transfer technology and thus they must implement only technology specific features of the corresponding technology.

When updating features which are related to information covered also in the C interface it might happen that the data the node map refers to changes unexpectedly. Therefore these values should not be cached in the nodemap but read every time from the module. This especially applies to features under a module selector.

### 7.1 Common

The common feature set is mandatory for all GenTL Producer implementations and used for all transport layer technologies.

#### 7.1.1 System Module

This is a description of all features which must be accessible in the System module: Port functions use the TL_HANDLE to access these features. The Port access for this module is mandatory.

Table 7-5: System module information features

|  Name | Interface | Access | Description  |
| --- | --- | --- | --- |
|  TLVendorName | IString | R | Name of the GenTL Producer vendor.  |
|  TLModelName | IString | R | Name of the GenTL Producer to distinguish different kinds of GenTL Producer implementations from one vendor.  |
|  TLID | IString | R | Unique ID identifying a GenTL Producer. For example the filename of the GenTL Producer implementation  |