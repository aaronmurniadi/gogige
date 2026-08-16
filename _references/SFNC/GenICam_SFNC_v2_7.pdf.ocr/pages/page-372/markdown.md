|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

OffsetY[Scan3dExtraction0] = 0;

ComponentSelector[Scan3dExtraction0] = Range;
ComponentEnable[Scan3dExtraction0][Range] = True;
PixelFormat[Scan3dExtraction0][Range] = Coord3D_C16
ComponentSelector[Scan3dExtraction0] = Reflectance;
ComponentEnable[Scan3dExtraction0][Reflectance] = True;
PixelFormat[Scan3dExtraction0][Reflectance] = Mono8;
ComponentSelector[Scan3dExtraction0] = Scatter;
ComponentEnable[Scan3dExtraction0][Scatter] = True;
PixelFormat[Scan3dExtraction0][Scatter] = Mono16;

// Sensor has 2000 pixels across, FOV range ~200 mm.
// Anchor system is "at edge of belt" Z towards camera.
// Invalid data tagged as 2047 in range map.
// Data is 16 bit unsigned integers.
// X axis has 0.5 mm per pixel in the rectified image, covering [-0.5,0.5] meter.
Scan3dCoordinateSelector = CoordinateA; // CoordinateA -> X
scaleA = Scan3dCoordinateScale[CoordinateA]; // e.g. 0.5
offsetA = Scan3dCoordinateOffset[CoordinateA]; // e.g. -500

// Coordinate B scaling
Scan3dCoordinateSelector = CoordinateB; //CoordinateB -> Y
scaleB = Scan3dCoordinateScale[CoordinateB]; // e.g. 0.1
offsetB = Scan3dCoordinateOffset[CoordinateB]; // e.g. -450

// Example : Z data has 0.1 mm per increment, no offset,
// 2047 is invalid data flag, 0-2000 valid data range.
Scan3dCoordinateSelector = CoordinateC; // CoordinateC -> Z
scaleC = Scan3dCoordinateScale[CoordinateC]; // e.g. 0.1
offsetC = Scan3dCoordinateOffset[CoordinateC]; // e.g. 0
usingInvalidFlagC = Scan3dInvalidDataFlag[CoordinateC]; // Invalid flag state.
invalidValueC = Scan3dInvalidDataValue[CoordinateC]; // e.g. 2047
bboxMinC = Scan3dAxisMin[CoordinateC]; // e.g. 0
bboxMaxC = Scan3dAxisMax[CoordinateC]; // e.g. 2000

Alternatively the information for scaling and offset is read from the chunk data for each image line in the RectifiedC_Linescan mode using the following pseudo code:

// Linescan 3D Camera "rectified".
// ***
// Basic device (No Region or Scan3dExtraction Selectors).
Scan3dOutputMode = RectifiedC_Linescan; // Sets B (Y) position given by encoder.
// Setup to use Chunk encoder data for Y scaling.
ChunkModeActive = True;
ChunkSelector = Image
ChunkEnable[ChunkSelector] = True;
ChunkSelector = Scan3dCoordinateScale;
ChunkEnable[ChunkSelector] = True;
ChunkSelector = Scan3dCoordinateOffset;
ChunkEnable[ChunkSelector] = True;
ChunkSelector = EncoderValue;