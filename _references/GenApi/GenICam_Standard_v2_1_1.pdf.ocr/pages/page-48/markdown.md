|  GENICAM |   | ![img-64.jpeg](img-64.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

## 3 Appendix

### 3.1 Endianess of GigE Vision Cameras

Because the GigE Vision standard provides two different schemes for the register access (READMEM and READREG) and because this fact was frequent source of confusion among different GenICam implementations in past, this section clarifies, how GenICam endianess should be implemented by GigE Vision based products.

For historical purposes, it defines two different kinds of behavior, each targeting different GenICam schema version:

- Behavior of products using GenICam schema version 1.1 and higher - this is the "correct" behavior, allowing full flexibility and no extra limitations.
- Behavior of products using GenICam schema version 1.0 - "legacy" attitude maintained for backward compatibility with schema 1.0 based products. This attitude has several limitations (or undefined behavior), especially for the little endian cameras.

Cameras providing XML files for different schema versions (eg. 1.1 and 1.0) are possible; the two attitudes differ only on side of the XML file. The behavior of the camera firmware (how registers are accessed) is defined by the GigE Vision standard and is thus independent on the GenICam.

While the behavior of products using schema 1.1 and newer is normative (all devices and applications must fully respect it), the part treating schema 1.0 is just a recommendation of the typical expected behavior promising best interoperability (it cannot be normative, because schema 1.0 devices and applications were deployed prior creation of this document).

Note that the following discussion targets only the GigE Vision cameras, it has no effect on other transport technologies or other GenApi uses.

#### 3.1.1 Behavior of products based on schema version 1.1 and newer

Cameras providing XML file based on schema version 1.1 or newer must implement endianess as follows:

- The <endianess> tags of all registers have to correspond with the real endianess of the camera, corresponding with the endianess reported in the DeviceMode bootstrap register.
- The port's <swapendianess> tag must not be used.</swapendianess></endianess>

It should also follow the requirements/recommendations listed in the GigE Vision specification (note that these are not mandatory requirements in the GigE Vision standard, just recommendations, but they become important when the camera is accessed through the GenICam interface):

- WRITEMEM should be implemented (if possible) when device is to be used through a GenICam interface.
- If READREG/WRITEREG is used to access strings, the camera must behave as if the string is composed of multiple 32-bit registers. This means that little endian camera must flip (reverse) each 4-character group, as expected. The same rules apply when accessing integers/floats of different size than 32-bit (particularly 64-bit and 16-bit ones).