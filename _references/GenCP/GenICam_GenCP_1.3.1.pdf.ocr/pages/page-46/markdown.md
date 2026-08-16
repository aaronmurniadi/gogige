|  GEN<i>CAM |   | ![img-49.jpeg](img-49.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

## 5. Bootstrap Register Map

### 5.1. Technology Agnostic Bootstrap Register Map

The Technology Agnostic Bootstrap Register Map (ABRM) uses the first 64 Kbytes of the register space. The table below shows the layout of the technology agnostic part of that bootstrap register map. This part also contains pointers to various other parts like the Manifest which provides access to the device GenICam files or the technology specific bootstrap registers.

### 5.2. String Registers

String registers not fully used are to be filled with 0. In case the full register is used, the terminating 0 can be omitted. The encoding of the content of a string register must match the Device Capability register.

### 5.3. Conditional Mandatory Registers

Conditional Mandatory (CM) registers are registers which may or may not be implemented depending on the Device Capability register. Access to a CM register which is indicated as being not available will return a GENCP_INVALID_ADDRESS status code.