|  GEN<i>CAM |   | ![img-42.jpeg](img-42.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|   | Mono2pMono4pMono8Mono10Mono10p...  |
| --- | --- |

Select the pixel format for which the information will be returned.

The pixel format selected must be one of the values present in the PixelFormat feature.

Possible values are:

• Mono1p: Mono 1 bit packed.
• Mono2p: Mono 2 bit packed.
• Mono4p: Mono 4 bit packed.
• Mono8: Mono 8 bit packed.
• Mono10: Mono 10 bit.
• Mono10p: Mono 10 bit packed.
• ...

Note: This feature must be a floating node and should always be available.

4.37 PixelFormatInfoID

|  Name | PixelFormatInfoID[PixelFormatInfoSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | -  |

Returns the value used by the streaming channels to identify the selected pixel format.

This value is generally equal to the standardized GenICam PFNC value for the selected PixelFormat.

To change the Pixel format of the data that will be sent by the device, the PixelFormat feature should be used.

Note: This feature must be a floating node and should always be available.

4.38 PixelCoding (Deprecated)

|  Name | PixelCoding  |
| --- | --- |