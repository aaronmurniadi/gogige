VariableFields (continued)

|  Width (bits) | Bit offset (lsb << x) | Description  |
| --- | --- | --- |
|  1 | 6 | **ComponentCount** If True, some Components might not be transmitted. The omitted Components must be the last ones.  |
|  1 | 7 | **ComponentInvalid** If True, the **Invalid** flag of certain Components of the Container might change.  |
|  8 | 8 | **Reserved** (set to 0).  |

|  6 | 26 | **Reserved** Reserved for future use (set to 0).  |
| --- | --- | --- |
|  8 | 32 | **DataSize** Maximum size of the entire Container's data section in bytes. Note: Does not include the size of the Container Descriptor.  |
|  8 | 40 | **DataOffset** Offset in bytes of the start of the entire Container's data section relative to the start of the Container's Descriptor.  |
|  4 | 48 | **DescriptorSize** Size of the Container Descriptor in bytes. This represents the cumulative size of the Container Header, all the Component Headers and their Part Headers.  |
|  4 | 52 | **ComponentCount** Number of Components in the entire Container.  |
|  Component Count x8 | 56 | **ComponentOffset[]** Array of the offsets in bytes of the start of each of the Component Headers relative to the start of the Container's Descriptor. The size of the array is **ComponentCount** x 8 bytes.  |

Table 2-1: GenDC Container Header Description