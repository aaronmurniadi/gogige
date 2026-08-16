|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Scan3dCoordinateOffset[Scan3dExtraction2] = 225;
Scan3dCoordinateScale[Scan3dExtraction2][CoordinateC] = -0.1;

// 3.c. Setup Processing Region functionality.
RegionSelector = Scan3dExtraction2;
RegionMode[Scan3dExtraction2] = On; // Transmit 3D Scan Extraction.
Width[Scan3dExtraction2] = 100; // Number of columns outbut buffer.
Height[Scan3dExtraction2] = 500; // Scans per buffer.
ComponentSelector[Scan3dExtraction2] = Range;
ComponentEnable[Scan3dExtraction2][Range] = True;
PixelFormat[Scan3dExtraction2][Range] = Coord3D_ABC32f;

The data output could in this case be formatted as:

|  Component | Part | Source | Region | data type | data format  |
| --- | --- | --- | --- | --- | --- |
|  Range | 0 | 1 | Scan3dExtraction 0 | 3D point cloud image | Coord3D_ABC32f  |
|  Range | 1 | 1 | Scan3dExtraction 1 | 3D point cloud image | Coord3D_ABC32f  |
|  Range | 2 | 1 | Scan3dExtraction 2 | 3D point cloud image | Coord3D_ABC32f  |
|  Chunk data | (3) | - | All components chunk data if enabled (not shown). | (Chunk data) | GenICam Chunk  |

### Calculating World Coordinates:

There are multiple methods on how to transform transmitted coordinates to world coordinates using the scale and offsets. The Scan3dOutputMode defines which method is applicable, if the mode is rectified, the X and/or Y coordinates are implicit and the world coordinates are calculated from the pixel coordinates.

In uncalibrated data transfer scaling and offset can also be used to allow meaningful scaling, for instance to allow mapping the result to the sensors coordinate system.

The pseudo code shows how to read and use scale and offset in a full 3D calibrated image using ChunkData. The SourceID and RegionID information for a component/part should be given by the TransportLayer, so that the selectors in the ChunkData can be set correctly.

// Reading chunk data from received buffer and scaling 3D image to world coordinates.
// For the Scan3dExtraction0 processing Region:
ChunkRegionSelector = Scan3dExtraction0;
CoordinateSystem = ChunkScan3dCoordinateSystem[Scan3dExtraction0]; // Example is for Cartesian
ScanMode = ChunkScan3dOutputMode[Scan3dExtraction0];

if (ScanMode == CalibratedABC_Grid)
{
    // Calculating World coordinates - Framescan.
    // ***
    // CoordinateA -> X
    ChunkScan3dCoordinateSelector = CoordinateA;