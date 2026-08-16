|  **GEN<i>CAM** |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

## 2 Summary of the Pixel Naming Convention

A pixel name is a text string composed of the following 5 fields, the last 3 having default values when they are not explicitly indicated.

|  Components & Location | # bits | [data type] | [packing] | [interface-specific]  |
| --- | --- | --- | --- | --- |

Figure 2-1 : Naming Convention Text Fields

Table 2-1 : Naming Convention Text Fields

|  Field | Description  |
| --- | --- |
|  **Components and Location** | Provides the list of components (ex: RGB, Y’CbCr, abstract A-B-C for 3D, …) and a reference to pixel location/sub-sampling if needed (ex: BayerRG, Y’CbCr422, …). In certain cases, an identifier might be used to differentiate between 2 similar color formats (Y’CbCr using ITU-R BT.601 vs ITU-R BT.709).  |
|  **# bits** | # of bits of each component  |
|  **Data type** (optional) | Data type indicator • *empty* or ‘**u**’: unsigned integer data • ‘**s**’: 2’s complement signed integer data • ‘**f**’: IEC 60559:1989 compliant floating point data  |
|  **Packing** (optional) | Packing style indicator showing how data is put into bytes and how to align them. • *empty*: unpacked data. Empty bits of each component must be padded with 0 to align to byte boundary. • ‘**p**’: packed data with no bit left in between components. • ‘**g**’: grouped data where least significant bits or most significant bits of the components are grouped in a separate byte. • ‘**c**’: cluster of single-component/monochrome pixels indicating the number of pixels to put together. This marker does not provide packing information per say. • ‘**a**’: an additional tag indicating the pixel is aligned to the given number of bits.  |
|  **Interface-specific** (optional) | This field is specific to the camera interface. It is the responsibility of the specific standard to define how to use this field. For instance, this field could be used to specify how data is ordered into data packets (sequencing of components in the packet) or on various image streams (ex: planar mode).  |