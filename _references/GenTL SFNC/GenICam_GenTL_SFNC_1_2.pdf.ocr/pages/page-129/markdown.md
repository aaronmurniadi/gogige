|   | BayerRG10 BayerRG10g12 BayerBG12 BayerBG12g BayerGB12 BayerGB12g BayerGR12 BayerGR12g BayerRG12 BayerRG12g BayerBG16 BayerGB16 BayerGR16 BayerRG16 Raw16 Raw8 Device-specific - GigE Vision Specific: Mono12Packed BayerGR10Packed BayerRG10Packed BayerGB10Packed BayerBG10Packed BayerGR12Packed BayerRG12Packed BayerGB12Packed BayerBG12Packed RGB10V1Packed BGR10V1Packed RGB12V1Packed  |
| --- | --- |

Format of the pixels provided by the buffer.

Note that the value list already follows the updated value list of the "PixelFormat" feature from GenICam SFNC 2.0, i.e., this feature does not exactly correspond to the

BUFFER_INFO_PIXELFORMAT command of DSGetBufferInfo function in the GenICam GenTL 1.3 standard. For multipart buffers this corresponds to

BUFFER_PART_INFO_DATA_FORMAT for PFNC formatted parts.

Note that only a subset of the possible pixel formats is listed here. The complete list of possible standard pixel formats and their detailed layout can be found in the "Pixel Format Naming