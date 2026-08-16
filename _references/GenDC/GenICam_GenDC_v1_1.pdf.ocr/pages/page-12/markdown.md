|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

[ R-001] A GenDC compliant product must use the Headers and flags as defined by this specification.

### 2.2.1 GenDC Container Header Layout

|  Signature = "GNDC" |   | Version Major | Version Minor | Version Sub Minor | Reserved  |
| --- | --- | --- | --- | --- | --- |
|  HeaderType | Flags | HeaderSize  |   |   |   |
|  Id  |   |   |   |   |   |
|  VariableFields | Reserved  |   |   |   |   |
|  DataSize  |   |   |   |   |   |
|  DataOffset  |   |   |   |   |   |
|  DescriptorSize |   | ComponentCount  |   |   |   |
|  ComponentOffset[ComponentCount]  |   |   |   |   |   |
|  ...  |   |   |   |   |   |

Figure 2-3: GenDC Container Header Layout