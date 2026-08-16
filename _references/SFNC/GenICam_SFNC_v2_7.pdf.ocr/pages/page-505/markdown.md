|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Category | TransportLayerControl  |
| --- | --- |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ImageSize PixelFormat PixelSize ExtendedPayload Device-specific  |

Selects the type of feature for which the locking behavior will be configured.

Possible values are:

- **ImageSize**: Controls if the images size is locked during the acquisition. ImageSize represents width and height for area-scan cameras; it only represents width for linescan cameras.
- **PixelFormat**: Controls if the pixel format (**PixelFormat**) is locked during an acquisition.
- **PixelSize**: Controls if the size of the pixel can change during an acquisition. When locked, the **PixelFormat** is allowed to change as long as **PixelSize** is not affected.
- **ExtendedPayload**: Controls if a device is allowed to switch between a payload and its extended version during an acquisition.

### 27.2.4 TLParamsLockedState

|  Name | TLParamsLockedState[TLParamsLockedSelector]  |
| --- | --- |
|  Category | TransportLayerControl  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | True False  |