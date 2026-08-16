|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Scan3dCoordinateOffset[Scan3dExtraction0] = 125;
Scan3dCoordinateScale[Scan3dExtraction0][CoordinateC] = 0.14;

// 1.c. Setup 3D Processing output Region and Components.
RegionSelector = Scan3dExtraction0;
RegionMode[Scan3dExtraction0] = On; // Transmit 3D Scan Extraction.
Width[Scan3dExtraction0] = 100; // Number of columns outbut buffer.
Height[Scan3dExtraction0] = 500; // Scans/Rows per output buffer.
ComponentSelector[Scan3dExtraction0] = Range;
ComponentEnable[Scan3dExtraction0][Range] = True;
PixelFormat[Scan3dExtraction0][Range] = Coord3D_ABC32f;

// 2. Dataflow: Region1 -> Scan3dExtraction1 -> Stream0
// ---

// 2.a. Setup Sensor Region output.
RegionSelector = Region1;
RegionMode[Region1] = On; // Transmit Sensor Region.
Height[Region1] = 256; // Number of rows on sensor.
Width[Region1] = 100; // Number of columns on sensor.
OffsetX[Region1] = 512; // Position on sensor.
OffsetY[Region1] = 25; // Position on sensor.

// 2.b. Setup 3D Processing Module source and parameters.
Scan3dExtractionSelector = Scan3dExtraction1;
Scan3dOutputMode[Scan3dExtraction1] = CalibratedC;
Scan3dExtractionSource[Scan3dExtraction1] = Region1;
Scan3dExtractionMethod[Scan3dExtraction1] = Default;
Scan3dCoordinateOffset[Scan3dExtraction1] = 125;
Scan3dCoordinateScale[Scan3dExtraction1][CoordinateC] = 0.14;

// 2.c. Setup Processing Region functionality
RegionSelector = Scan3dExtraction1;
RegionMode[Scan3dExtraction1] = On; // Transmit 3D Scan Extraction.
Width[Scan3dExtraction1] = 100; // Number of columns outbut buffer.
Height[Scan3dExtraction1] = 500; // Scans per buffer.
ComponentSelector[Scan3dExtraction1] = Range;
ComponentEnable[Scan3dExtraction1][Range] = True;
PixelFormat[Scan3dExtraction1][Range] = Coord3D_ABC32f;

// 3. Dataflow: Region2 -> Scan3dExtraction2 -> Stream0
// ---

// 3.a. Setup Sensor Region output.
RegionSelector = Region2;
RegionMode[Region2] = On; // Transmit Sensor Region.
Height[Region2] = 256; // Number of rows on sensor.
Width[Region2] = 200; // Number of columns on sensor.
OffsetX[Region2] = 50; // position on sensor.
OffsetY[Region2] = 400; // position on sensor.

// 3.b. Setup 3D Processing Module source and parameters.
Scan3dExtractionSelector = Scan3dExtraction2;
Scan3dOutputMode[Scan3dExtraction2] = CalibratedC;
Scan3dExtractionSource[Scan3dExtraction2] = Region2;
Scan3dExtractionMethod[Scan3dExtraction2] = Default;