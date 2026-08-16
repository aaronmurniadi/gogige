|   | RGBA8Packed (Deprecated, use RGBa8) BGRA8Packed (Deprecated, use BGRa8) RGB10Packed (Deprecated, use RGB10) BGR10Packed (Deprecated, use BGR10) RGB12Packed (Deprecated, use RGB12) BGR12Packed (Deprecated, use BGR12) RGB16Packed (Deprecated, use RGB16) BGR16Packed (Deprecated, use BGR16) RGB10V2Packed (Deprecated, use RGB10p32) BGR10V2Packed (Deprecated, use BGR10p32) RGB565Packed (Deprecated, use RGB565p) BGR565Packed (Deprecated, use BGR565p) YUV411Packed (Deprecated, use YUV411_8_UYYVYY) YUV422Packed (Deprecated, use YUV422_8_UYVY) YUV444Packed (Deprecated, use YUV8_UYV) YUYVPacked (Deprecated, use YUV422_8) RGB8Planar (Deprecated, use RGB8_Planar) RGB10Planar (Deprecated, use RGB10_Planar) RGB12Planar (Deprecated, use RGB12_Planar) RGB16Planar (Deprecated, use RGB16_Planar)  |
| --- | --- |

Format of the pixels provided by the device. It represents all the information provided by PixelSize, PixelColorFilter combined in a single feature.

Note that only a subset of the possible pixel formats is listed in this document. The complete list of currently standardized pixel formats and their assigned identifier value can be found in the "GenICam Pixel Format Values" document or in the "Reference Header file for PFNC" on the GenICam download page on the EMVA web site. See also the "GenICam Pixel Format Naming Convention" document for information on the detailed layout of those standard formats.

This feature is generally mandatory for transmitters and transceivers of most Transport Layers. You can refer to the appropriate transport layer specification for additional information.

Note: Transport Layer standards using frame grabbers on the receiver side may code the pixels differently than the PFNC on the link (to save bandwidth for example). In that case, the standard will define how the data is coded for each supported pixel format, and will define the requirements for the frame grabber to convert the data into a PFNC compliant pixel format.

Possible values are:

- Mono1p: Mono 1 bit packed.
- Mono2p: Mono 2 bit packed.
- Mono4p: Mono 4 bit packed.
- Mono8: Mono 8 bit packed.
- Mono8s: Mono 1 bit signed.
- Mono10: Mono 10 bit.
- Mono10p: Mono 10 bit packed.
- Mono12: Mono 12 bit.