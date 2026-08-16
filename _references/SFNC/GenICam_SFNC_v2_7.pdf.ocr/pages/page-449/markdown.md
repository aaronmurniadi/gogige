|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

...
// Enable the Range and Confidence Components.
ComponentSelector = Range;
ComponentEnable = True;
ComponentSelector = Confidence;
ComponentEnable = True;

// Enable the Image and PixelFormat chunks.
ChunkModeActive = True;
ChunkSelector = Image;
ChunkEnable = True;
ChunkSelector = PixelFormat;
ChunkEnable = True;

// Start the acquisition of the multi-component buffer with its chunk data.
AcquisitionStart();
...

# Reception side:

// At reception of a multi-component payload, the components' pixel format are retrieved using:
ChunkComponentSelector = Range;
RangePixelFormat = ChunkPixelFormat[Range];
ComponentSelector = Confidence;
ConfidencePixelFormat = ChunkPixelFormat[Confidence];

# ChunkScanLineSelector:

Linescan cameras can have features that vary on per line basis. ExposureTime and EncoderValue are examples of such features that can vary from line to line. In order to handle properly those features in linescan mode, their value for each scan line must be included in the Chunk Data. ChunkScanLineSelector can then be used to retrieve each value in the received chunk.

# Chunk Features description:

The chunk features are described below.

Note that the naming scheme to access the data of the chunk associated with a feature Name is ChunkName and any SFNC feature can have a chunk counterpart when this applies.

# 24.1 ChunkDataControl

|  Name | ChunkDataControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |