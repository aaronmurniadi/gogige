|  GEN<i>CAM |   | ![img-69.jpeg](img-69.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 5.4.22. Device Software Interface Version

The Device Software Interface Version references a certain version of the publicly available software interface of the device. The content of the register should change to a new value (not used before) whenever any of this changes:

- implemented communication protocol
- publicly available register map (all registers referenced by the XML and the bootstrap)
- user accessible camera functionality
- the GenApi XML.

The semantics of the string are vendor specific. The standard only requires that the string changes if any of the above listed components change.

If this register is supported the according bit in the Device Capability register needs to be set to 1.

The Device Software Interface Version may or may not indicate some device internal changes but that is not the primary objective.

|  Offset | Hex 210  |
| --- | --- |
|  Length | 64  |
|  Access Type | R  |
|  Support | CM(intended to make M in the next major release of the this standard)  |
|  Data Type | String  |
|  Factory Default | Device specific  |

It is intended to make the Device Software Interface Version register mandatory in the next major release of this standard.

### 5.5. Generic Tables

#### 5.5.1. Manifest

The manifest provides a way to store multiple GenICam-related files in the device. These GenICam files may be available in different versions, in various formats or comply to different versions of the GenICam schema. The manifest table contains a list of Manifest Entries.