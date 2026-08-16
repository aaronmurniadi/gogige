|  GEN<I>CAM |   | ![img-95.jpeg](img-95.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

version number and a minor version number in a notation “x.y” with ‘x’ being the major version number and ‘y’ being the minor version number.

• Major Version Numbers

Different major version numbers indicate major additions to the interface and/or breaking changes. This means for example a removal of functions or a complete new feature set. The newer interface is therefore not backward compatible.

• Minor Version Numbers

Changes in the minor version number of the software interface may indicate new functionality and clarifications in the text describing the interface. If only the minor version changes the interface stays backward compatible.

Changing feature names without functionl change is also allowed in minor releases.

When developing a GenTL Consumer that should be compatible with a widest range of GenTL Producer versions, special care might be required to consider these differences.

When using an enumeration unknown to the GenTL Producer, the function getting that value as a parameter would return an appropriate error code. For example when querying an unknown info command, the GenTL Producer would return GC_ERR_NOT_IMPLEMENTED.

When trying to use a GenTL interface function unknown to the GenTL Producer, the function implementation will be simply missing in the GenTL Producer's binary. For the functions that are not universally available in all GenTL specification versions, the Consumer should check their presence in the GenTL Producer's interface at load time and if possible, consider a suitable fallback behaviour for GenTL Producers not implementing that function.

### 6.2 Used Data Types

To have a defined stack layout certain data types have a primitive data type as its base.

#### 6.2.1.1 GC_ERROR

The return value of all functions is a 32 bit signed integer value.

#### 6.2.1.2 Handles

All handles like TL_HANDLE or PORT_HANDLE are void*. The size is platform dependent (e.g., 32 bit on 32 bit platforms).

#### 6.2.1.3 Enumerations

All enumerations are of type enum. In order to allow implementation specific extensions all enums are set to a specific 32 bit integer value. On platforms/compilers where this is not the case a primitive data type with a matching size has to be used.

#### 6.2.1.4 Buffers and C Strings

Buffers are normally typed as void* if arbitrary data is accessed. Specialized buffers like C strings are by default ASCII encoded and a char* is used unless reported different through the type information provided by the info functions (for example IFGetInfo). A char is